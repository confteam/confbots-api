package main

import (
	"context"
	"fmt"
	"time"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/config"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/infrastructure/repository"
	"github.com/confteam/confbots-api/internal/logger"
)

func main() {
	// TODO:
	// usecase
	// router
	// server

	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	pool, err := repository.NewPgxPool(cfg.DBConfig)
	if err != nil {
		log.Error("cannot connect to db", "error", err)
		return
	}
	queries := db.New(pool)
	botRepo := repository.NewBotPostgresRepository(queries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	bot, err := botRepo.CreateIfNotExists(ctx, 2089144368, entities.BotTypeTakes)
	if err != nil {
		log.Error("cannot create bot", "error", err)
		return
	}
	fmt.Println(bot)
}
