package director

import (
	"net"
	"path/filepath"
	"sync"

	"github.com/grogdb/grogdb/internal/cluster"
	"github.com/rs/zerolog"
)

func NewDirector(logger *zerolog.Logger, cluster cluster.Cluster, dataDir string, raftAddr, raftAdvertiseAddr *net.TCPAddr, raftDebug bool) Director {
	l := logger.With().Str("component", "director").Logger()
	return &director{
		cluster:           cluster,
		dataDir:           filepath.Join(dataDir, "director"),
		logger:            &l,
		raftAddr:          raftAddr,
		raftAdvertiseAddr: raftAdvertiseAddr,
		raftDebug:         raftDebug,
	}
}

type Director interface {
	Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
}

type director struct {
	cluster           cluster.Cluster
	dataDir           string
	logger            *zerolog.Logger
	raftAddr          *net.TCPAddr
	raftAdvertiseAddr *net.TCPAddr
	raftDebug         bool
}

func (d *director) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	stopWg.Add(1)
	go d.startRaft(stopWg, shutdownCh, errCh)
}

func (d *director) startRaft(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()
}
