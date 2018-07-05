package main

import (
	"fmt"
	"os"
	"github.com/grogdb/grog/internal/bootstrap"
)

var (
	Build = "dev"
	Version = "X.X.X"
)

func main() {
	if err := bootstrap.Execute(Build, Version); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func bootstrap(svcLogger *zerolog.Logger, mongoURL string, httpAddr string, tokenKey string) error {
// 	svcLogger.Debug().Msgf("connecting to MongoDB: %s", mongoURL)
//
// 	mongoClient, err := mongo.NewClient(mongoURL)
// 	if err != nil {
// 		return err
// 	}
//
// 	if err = mongoClient.Connect(context.Background()); err != nil {
// 		return err
// 	}
// 	defer mongoClient.Disconnect(context.Background())
//
// 	svcLogger.Debug().Msgf("configuring service...")
// 	paydayRepo := mongodb.NewMongoRepository(mongoClient.Database("payday"))
// 	paydayService := service.NewService(paydayRepo)
//
// 	svcLogger.Debug().Msgf("configuring HTTP transport...")
// 	r := chi.NewRouter()
// 	r.Use(cors.CORS)
// 	r.Use(requestid.HTTPRequestID)
// 	r.Use(logger.HTTPRequestLogger(svcLogger))
// 	r.Use(middleware.Timeout(10 * time.Second))
// 	r.Use(gatekeeper.AuthHTTP(tokenKey))
// 	r.Use(middleware.Recoverer)
//
// 	httpHandler := handler.NewHandlerHTTP(svcLogger, paydayService)
// 	httpHandler.ConfigureRouter(r)
//
// 	svcLogger.Info().Msgf("starting HTTP transport on %s", httpAddr)
//
// 	server := &http.Server{Addr: httpAddr, Handler: r}
// 	serverErrCh := make(chan error, 1)
// 	defer close(serverErrCh)
//
// 	go func() {
// 		if err := server.ListenAndServe(); err != nil {
// 			serverErrCh <- err
// 		}
// 	}()
//
// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
//
// 	select {
// 	case <-stop:
// 		// wait for signal
// 	case err := <-serverErrCh:
// 		return err
// 	}
//
// 	svcLogger.Info().Msgf("stopping HTTP transport on %s", httpAddr)
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	if err := server.Shutdown(ctx); err != nil {
// 		return err
// 	}
//
// 	<-ctx.Done()
// 	return nil
// }
