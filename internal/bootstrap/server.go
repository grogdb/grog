package bootstrap

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/grafodb/grafodb/internal/cluster"
	"github.com/grafodb/grafodb/internal/director"
	"github.com/grafodb/grafodb/internal/graph"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	BootstrapExpect   int
	ClusterRPCAddr    *net.TCPAddr
	DataDir           string
	DirectorEnabled   bool
	GraphEnabled      bool
	GraphHTTPAddr     *net.TCPAddr
	GraphRPCAddr      *net.TCPAddr
	JoinAddrs         []string
	SerfAddr          *net.TCPAddr
	SerfAdvertiseAddr *net.TCPAddr
	SerfEncryptionKey string
	SerfDebug         bool
	RaftAddr          *net.TCPAddr
	RaftAdvertiseAddr *net.TCPAddr
	RaftDebug         bool
}

func NewServer(logger *zerolog.Logger, cfg *ServerConfig) Server {
	return &server{
		config:     cfg,
		logger:     logger,
		shutdownCh: make(chan struct{}, 1),
	}
}

type Server interface {
	Start() error
	Shutdown()
}

type server struct {
	config       *ServerConfig
	cluster      cluster.Cluster
	graph        graph.Graph
	director     director.Director
	logger       *zerolog.Logger
	shutdownLock sync.Mutex
	shutdownCh   chan struct{}
	shutdown     bool
}

func (s *server) Start() error {
	var stopWg sync.WaitGroup
	startErrCh := make(chan error, 1)

	cfg := s.config

	c, err := cluster.NewCluster(s.logger, cfg.DataDir, cfg.ClusterRPCAddr, cfg.SerfAddr, cfg.SerfAdvertiseAddr, cfg.SerfDebug)
	if err != nil {
		return err
	}
	s.cluster = c
	s.cluster.Serve(&stopWg, s.shutdownCh, startErrCh)

	if cfg.DirectorEnabled {
		s.director = director.NewDirector(s.logger, s.cluster, cfg.DataDir, cfg.RaftAddr, cfg.RaftAdvertiseAddr, cfg.RaftDebug)
		s.director.Serve(&stopWg, s.shutdownCh, startErrCh)
	} else {
		s.logger.Info().Msg("Director not enabled for this node")
	}

	if cfg.GraphEnabled {
		s.graph = graph.NewGraph(s.logger, s.cluster, cfg.DataDir, cfg.GraphHTTPAddr, cfg.GraphRPCAddr)
		s.graph.Serve(&stopWg, s.shutdownCh, startErrCh)
	} else {
		s.logger.Info().Msg("Graph not enabled for this node")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	var startErr error
	select {
	case err := <-startErrCh:
		startErr = err
	case <-quit:
		log.Info().Msg("Initializing server shutdown")
	}
	s.Shutdown()

	s.logger.Debug().Msg("Waiting for shutdown to complete...")
	stopWg.Wait()
	s.logger.Info().Msg("Shutdown completed")

	return startErr
}

func (s *server) Shutdown() {
	s.shutdownLock.Lock()
	defer s.shutdownLock.Unlock()

	if s.shutdown {
		return
	}
	s.shutdown = true
	close(s.shutdownCh)
}
