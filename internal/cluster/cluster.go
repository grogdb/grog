package cluster

import (
	"fmt"
	"io/ioutil"
	stdlog "log"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/grafodb/grafodb/internal/util/fs"
	"github.com/grafodb/grafodb/internal/util/random"
	"github.com/hashicorp/serf/serf"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

const (
	tagIsDirector = "is_director"
	tagIsGraph    = "is_graph"
)

type Config struct {
	ClusterRPCAddr    *net.TCPAddr
	DataDir           string
	IsDirector        bool
	IsGraph           bool
	GraphHTTPAddr     *net.TCPAddr
	GraphRPCAddr      *net.TCPAddr
	JoinAddrs         []string
	SerfAddr          *net.TCPAddr
	SerfAdvertiseAddr *net.TCPAddr
	SerfEncryptionKey string
	SerfDebug         bool
	RaftAdvertiseAddr *net.TCPAddr
}

func NewCluster(logger *zerolog.Logger, dataDir string, rpcAddr, serfAddr, serfAdvertiseAddr *net.TCPAddr, serfDebug bool) (Cluster, error) {
	log := logger.With().Str("component", "cluster").Logger()

	clusterDataDir := filepath.Join(dataDir, "cluster")
	serfDataDir := filepath.Join(clusterDataDir, "serf")

	if err := fs.EnsureDir(clusterDataDir); err != nil {
		return nil, fmt.Errorf("error creating cluster data dir: %s", clusterDataDir)
	}

	if err := fs.EnsureDir(serfDataDir); err != nil {
		return nil, fmt.Errorf("error creating serf data dir: %s", serfDataDir)
	}

	return &cluster{
		clusterDataDir:    clusterDataDir,
		eventCh:           make(chan serf.Event, 256),
		logger:            &log,
		members:           make(map[string]MemberInfo),
		rpcAddr:           rpcAddr,
		serfAddr:          serfAddr,
		serfAdvertiseAddr: serfAdvertiseAddr,
		serfDataDir:       serfDataDir,
		serfDebug:         serfDebug,
	}, nil
}

type Cluster interface {
	Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
}

type cluster struct {
	clusterDataDir    string
	eventCh           chan serf.Event
	logger            *zerolog.Logger
	memberID          string
	members           map[string]MemberInfo
	membersLock       sync.RWMutex
	rpcAddr           *net.TCPAddr
	serf              *serf.Serf
	serfAddr          *net.TCPAddr
	serfAdvertiseAddr *net.TCPAddr
	serfDataDir       string
	serfDebug         bool
}

func (c *cluster) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	stopWg.Add(2)
	go c.startServerRPC(stopWg, shutdownCh, errCh)
	go c.startSerf(stopWg, shutdownCh, errCh)
}

func (c *cluster) startServerRPC(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	grpcServerAddr := c.rpcAddr.String()
	grpcServer := grpc.NewServer()

	// FIXME: Bind RPC
	// s.cluster.BindRPC(grpcServer)

	startErrCh := make(chan error, 1)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		c.logger.Debug().Msgf("Starting cluster RPC server on %v", grpcServerAddr)
		listener, err := net.Listen("tcp", grpcServerAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		c.logger.Info().Msgf("Started cluster RPC server on %v", grpcServerAddr)
		startErrCh <- grpcServer.Serve(listener)
	}

	stopFunc := func() {
		c.logger.Debug().Msg("Stopping cluster RPC server...")
		grpcServer.GracefulStop()
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		close(doneCh)
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	c.logger.Info().Msg("Stopped cluster RPC server")
}

func (c *cluster) startSerf(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	if err := c.ensureMemberID(); err != nil {
		errCh <- fmt.Errorf("error ensuring cluster member id")
		return
	}

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		c.logger.Info().Msgf("Starting serf server on %v", c.serfAddr.String())

		serfConf := serf.DefaultConfig()
		serfConf.Init()
		// serfConf.Tags["role"] = m.role
		// serfConf.Tags["serf_port"] = fmt.Sprintf("%d", serfAddr.Port)
		// serfConf.Tags["rpc_port"] = fmt.Sprintf("%d", m.rpcPort)

		serfConf.EventCh = c.eventCh
		serfConf.NodeName = c.memberID
		serfConf.SnapshotPath = filepath.Join(c.serfDataDir, "serf.snapshot")
		serfConf.MemberlistConfig.BindAddr = c.serfAddr.IP.String()
		serfConf.MemberlistConfig.BindPort = c.serfAddr.Port
		serfConf.MemberlistConfig.AdvertiseAddr = c.serfAdvertiseAddr.IP.String()
		serfConf.MemberlistConfig.AdvertisePort = c.serfAdvertiseAddr.Port
		serfConf.MemberlistConfig.LogOutput = ioutil.Discard

		if c.serfDebug {
			serfConf.Logger = stdlog.New(c.logger, "", 0)
		} else {
			serfConf.LogOutput = ioutil.Discard
		}

		serfServer, err := serf.Create(serfConf)
		if err != nil {
			errCh <- err
			return
		}
		c.serf = serfServer

		go c.monitorClusterEvents(shutdownCh)
	}

	stopFunc := func() {
		c.logger.Debug().Msg("Stopping serf server...")
		if c.serf != nil {
			// TODO: Shall we leave the cluster on shutdown?
			c.serf.Leave()
			if err := c.serf.Shutdown(); err != nil {
				errCh <- err
			}
		}
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		stopFunc()
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	c.logger.Info().Msg("Stopped serf server")
}

func (c *cluster) ensureMemberID() error {
	if c.memberID != "" {
		return nil
	}

	memberIDFilePath := filepath.Join(c.serfDataDir, "node_id")

	if _, err := os.Stat(memberIDFilePath); err == nil {
		data, err := ioutil.ReadFile(memberIDFilePath)
		if err != nil {
			return err
		}
		c.memberID = string(data)
	} else {
		memberID := fmt.Sprintf("node-%s", random.HexString(8))
		if err := ioutil.WriteFile(memberIDFilePath, []byte(memberID), 0644); err != nil {
			return err
		}
		c.memberID = memberID
	}
	return nil
}

func (c *cluster) monitorClusterEvents(shutdownCh <-chan struct{}) {
	for {
		select {
		case e := <-c.eventCh:
			switch e.EventType() {
			case serf.EventMemberJoin:
				c.addMember(e.(serf.MemberEvent))
			case serf.EventMemberLeave, serf.EventMemberFailed:
				c.removeMember(e.(serf.MemberEvent))
			case serf.EventMemberUpdate:
				c.updateMember(e.(serf.MemberEvent))
				//case serf.EventMemberReap:
				//	//p.localMemberEvent(e.(serf.MemberEvent))
				//case serf.EventMemberUpdate, serf.EventUser, serf.EventQuery:
				//	// Ignore
			default:
				c.logger.Warn().Msgf("unhandled serf event: %#v", e)
			}
		case <-shutdownCh:
			return
		}
	}
}

func (c *cluster) addMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		ok, info := c.createMemberInfo(member)
		if !ok {
			c.logger.Warn().Msg("Attempt to add a member with an unknown role to the pool")
			continue
		}

		c.membersLock.Lock()
		if _, exists := c.members[info.ID]; !exists {
			c.logger.Info().Msgf("Adding member to pool: %s [%s]", info.ID, info.Addr)
			c.members[info.ID] = *info
		}
		c.membersLock.Unlock()
	}
}

func (c *cluster) removeMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		ok, info := c.createMemberInfo(member)
		if !ok {
			c.logger.Warn().Msg("Attempt to remove a member with an unknown role to the pool")
			continue
		}

		c.membersLock.Lock()
		if _, exists := c.members[info.ID]; exists {
			c.logger.Info().Msgf("Removing member from pool: %s [%s]", info.ID, info.Addr)
			delete(c.members, info.ID)
		}
		c.membersLock.Unlock()
	}
}

func (c *cluster) updateMember(ev serf.MemberEvent) {
	for _, member := range ev.Members {
		ok, info := c.createMemberInfo(member)
		if !ok {
			c.logger.Warn().Msg("Attempt to update a member with an unknown role to the pool")
			continue
		}

		c.membersLock.Lock()
		if _, exists := c.members[info.ID]; exists {
			c.logger.Info().Msgf("Updating member in pool: %s [%s]", info.ID, info.Addr)
			c.members[info.ID] = *info
		}
		c.membersLock.Unlock()
	}
}

func (c *cluster) createMemberInfo(member serf.Member) (bool, *MemberInfo) {
	// role := member.Tags["role"]
	// if role != RoleController && role != RoleWorker {
	// 	return false, nil
	// }
	//
	// rpcPort, err := strconv.Atoi(member.Tags["rpc_port"])
	// if err != nil {
	// 	return false, nil
	// }
	//
	// serfPort, err := strconv.Atoi(member.Tags["serf_port"])
	// if err != nil {
	// 	return false, nil
	// }
	//
	// parts := &MemberInfo{
	// 	ID:       member.Name,
	// 	Addr:     member.Addr.String(),
	// 	Role:     role,
	// 	RpcPort:  rpcPort,
	// 	SerfPort: serfPort,
	// 	Status:   member.Status.String(),
	// }
	// return true, parts
	return false, nil
}

/*
func (c *controller) startClusterManager(wg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer wg.Done()



	manager := cluster.NewClusterManager(&cluster.Config{
		AdvertiseIP:   c.config.AdvertiseIP,
		ClusterAddr:   c.config.ClusterAddr,
		Role:          cluster.RoleController,
		DataDir:       c.config.DataDir,
		EnableSerfLog: c.config.VerboseMode,
		RpcPort:       c.config.RpcAddr.Port,
	})
	c.clusterManager = manager

c
}
*/
