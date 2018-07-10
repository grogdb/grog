package graph

import (
	"context"
	"net"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/grogdb/grogdb/internal/cluster"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func NewGraph(logger *zerolog.Logger, cluster cluster.Cluster, dataDir string, httpAddr, rpcAddr *net.TCPAddr) Graph {
	l := logger.With().Str("component", "graph").Logger()
	return &graph{
		cluster:  cluster,
		dataDir:  filepath.Join(dataDir, "graph"),
		httpAddr: httpAddr,
		logger:   &l,
		rpcAddr:  rpcAddr,
	}
}

type Graph interface {
	Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error)
}

type graph struct {
	cluster  cluster.Cluster
	dataDir  string
	httpAddr *net.TCPAddr
	logger   *zerolog.Logger
	rpcAddr  *net.TCPAddr
}

func (g *graph) Serve(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	stopWg.Add(2)
	go g.startServerRPC(stopWg, shutdownCh, errCh)
	go g.startServerHTTP(stopWg, shutdownCh, errCh)
}

func (g *graph) startServerHTTP(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	httpServerAddr := g.httpAddr.String()
	httpServer := http.Server{
		Addr:    httpServerAddr,
		Handler: g.newRouterHTTP(),
	}

	startErrCh := make(chan error, 1)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		g.logger.Debug().Msgf("Starting graph HTTP server on %v", httpServerAddr)
		listener, err := net.Listen("tcp", httpServerAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		g.logger.Info().Msgf("Started graph HTTP server on %v", httpServerAddr)
		startErrCh <- httpServer.Serve(listener)
	}

	stopFunc := func() {
		g.logger.Debug().Msg("Stopping graph HTTP server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		httpServer.Shutdown(ctx)
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	g.logger.Info().Msg("Stopped graph HTTP server")
}

func (g *graph) startServerRPC(stopWg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer stopWg.Done()

	grpcServerAddr := g.rpcAddr.String()
	grpcServer := grpc.NewServer()

	// FIXME: bind RPC
	// s.graph.BindRPC(grpcServer)

	startErrCh := make(chan error, 1)
	doneCh := make(chan struct{}, 1)

	startFunc := func() {
		g.logger.Debug().Msgf("Starting graph RPC server on %v", grpcServerAddr)
		listener, err := net.Listen("tcp", grpcServerAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		g.logger.Info().Msgf("Started graph RPC server on %v", grpcServerAddr)
		startErrCh <- grpcServer.Serve(listener)
	}

	stopFunc := func() {
		g.logger.Debug().Msg("Stopping graph RPC server...")
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
	g.logger.Info().Msg("Stopped graph RPC server")
}

func (g *graph) newRouterHTTP() chi.Router {
	router := chi.NewRouter()
	// FIXME: bind HTTP
	return router
}
