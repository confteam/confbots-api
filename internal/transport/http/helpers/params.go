package helpers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/go-chi/chi/v5"
)

func ParseURLParam(w http.ResponseWriter, r *http.Request, log *slog.Logger, name string) (int, bool) {
	raw := chi.URLParam(r, name)
	id, err := strconv.Atoi(raw)
	if err != nil {
		log.Warn("invalid URL param", slog.String("param", name), slog.String("value", raw))
		EncodeJSON(w, r, http.StatusBadRequest, response.Error("invalid path parameter: "+name))
		return 0, false
	}
	return id, true
}

func ParseURLParamStr(w http.ResponseWriter, r *http.Request, log *slog.Logger, name string) (string, bool) {
	raw := chi.URLParam(r, name)
	if raw == "" {
		log.Warn("invalid URL param", slog.String("param", name), slog.String("value", raw))
		EncodeJSON(w, r, http.StatusBadRequest, response.Error("invalid path parameter: "+name))
		return raw, false
	}
	return raw, true
}

func ParseQuery(w http.ResponseWriter, r *http.Request, log *slog.Logger, name string, required bool) (int, bool) {
	raw := r.URL.Query().Get(name)

	if raw == "" {
		if required {
			log.Warn("missing query parameter", slog.String("param", name))
			EncodeJSON(w, r, http.StatusBadRequest, response.Error(fmt.Sprintf("%s is required", name)))
		}
		return 0, false
	}

	val, err := strconv.Atoi(raw)
	if err != nil {
		log.Warn("invalid query parameter", slog.String("param", name), slog.String("value", raw))
		EncodeJSON(w, r, http.StatusUnprocessableEntity, response.Error(fmt.Sprintf("failed to convert %s", name)))
		return 0, false
	}

	return val, true
}
