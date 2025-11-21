package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/confteam/confbots-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const op = "infrastructure.repository.NewPgxConn"

func NewPgxPool(dbCfg config.DBConfig, log *slog.Logger) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbCfg.Name,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
	))
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	log.Info("connected to db",
		slog.String("name", dbCfg.Name),
		slog.String("user", dbCfg.User),
		slog.String("host", dbCfg.Host),
		slog.String("port", dbCfg.Port),
		slog.String("db_name", dbCfg.DBName),
	)

	return pool, nil
}
