package helpers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/domain"
	"github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

func HandleError(w http.ResponseWriter, r *http.Request, log *slog.Logger, err error) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		log.Warn("User not found", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusNotFound, response.Error(err.Error()))

	case errors.Is(err, domain.ErrChannelNotFound):
		log.Warn("Channel not found", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusNotFound, response.Error(err.Error()))

	case errors.Is(err, domain.ErrBotNotFound):
		log.Warn("Bot not found", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusNotFound, response.Error(err.Error()))

	case errors.Is(err, domain.ErrUserChannelNotFound):
		log.Warn("UserChannel not found", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusNotFound, response.Error(err.Error()))

	case errors.Is(err, domain.ErrTakeNotFound):
		log.Warn("Take not found", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusNotFound, response.Error(err.Error()))

	case errors.Is(err, domain.ErrReplyNotFound):
		log.Warn("Reply not found", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusNotFound, response.Error(err.Error()))

	default:
		log.Error("internal server error", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusInternalServerError, response.Error("internal server error"))
	}
}
