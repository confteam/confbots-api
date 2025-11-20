package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/config"
	"github.com/confteam/confbots-api/internal/infrastructure/repository"
	"github.com/confteam/confbots-api/internal/logger"
	"github.com/confteam/confbots-api/internal/transport/http/handler"
	"github.com/confteam/confbots-api/internal/transport/http/server"
	"github.com/confteam/confbots-api/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	pool, err := repository.NewPgxPool(cfg.DBConfig)
	if err != nil {
		log.Error("cannot connect to db", "error", err)
		return
	}

	queries := db.New(pool)

	botRepo := repository.NewBotPostgresRepository(queries)
	channelRepo := repository.NewChannelPostgresRepository(queries)
	userRepo := repository.NewUserPostgresRepository(queries)

	botUseCase := usecase.NewBotUseCase(botRepo)
	channelUseCase := usecase.NewChannelUseCase(channelRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)

	botHandler := handler.NewBotHandler(botUseCase, log)
	channelHandler := handler.NewChannelHandler(channelUseCase, log)
	userHandler := handler.NewUserHandler(userUseCase, log)

	srv := server.NewServer(cfg.HTTPServer, log, func(r chi.Router) {
		botHandler.RegisterRoutes(r)
		channelHandler.RegisterRoutes(r)
		userHandler.RegisterRoutes(r)
	})

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.Start(); err != nil {
			log.Error("can't start server", "error", err)
			return
		}
	}()

	<-stop
	log.Info("recieved shutdown signal")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to shutdown server", "error", err)
		return
	}

	log.Info("server exited gracefully")
}
