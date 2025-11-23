package helpers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func ReqLogger(base *slog.Logger, r *http.Request, op string) *slog.Logger {
	return base.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
}
