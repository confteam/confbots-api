package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/config"
	mwLogger "github.com/confteam/confbots-api/internal/transport/http/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	srv *http.Server
	log *slog.Logger
}

func NewServer(cfg config.HTTPServer, log *slog.Logger, registerRoutes func(r chi.Router)) *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(mwLogger.New(log))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	registerRoutes(r)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.Timeout,
	}

	return &Server{
		srv: srv,
		log: log,
	}
}

const pkg = "transport.http.server.Server"

func (s *Server) Start() error {
	const op = pkg + "Start"
	s.log.Info("starting server", "address", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	const op = pkg + ".Shutdown"
	s.log.Info("shuttind down server")
	err := s.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}
