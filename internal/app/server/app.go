// Package app configures and runs application.
package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	config "github.com/dimk00z/GophKeeper/config/server"
	v1 "github.com/dimk00z/GophKeeper/internal/controller/http/v1"
	"github.com/dimk00z/GophKeeper/internal/usecase"
	"github.com/dimk00z/GophKeeper/internal/usecase/repo"
	"github.com/dimk00z/GophKeeper/internal/usecase/webapi"
	"github.com/dimk00z/GophKeeper/pkg/httpserver"
	"github.com/dimk00z/GophKeeper/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	gophKeeperRepo := repo.New(cfg.PG.URL, l)
	gophKeeperRepo.Migrate()

	// Use case
	GophKeeperUseCase := usecase.New(
		gophKeeperRepo,
		webapi.New(),
	)
	var err error
	// RabbitMQ RPC Server
	// rmqRouter := amqprpc.NewRouter(GophKeeperUseCase)

	// rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	// }

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, GophKeeperUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
