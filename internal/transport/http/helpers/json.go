package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func EncodeJSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	w.WriteHeader(status)
	render.JSON(w, r, v)
}

func DecodeJSON[T any](w http.ResponseWriter, r *http.Request, log *slog.Logger, dst *T) bool {
	err := render.DecodeJSON(r.Body, dst)
	if err == nil {
		return true
	}

	var syntaxErr *json.SyntaxError
	var unmarshalErr *json.UnmarshalTypeError

	switch {
	case errors.Is(err, io.EOF):
		log.Warn("empty request body")
		EncodeJSON(w, r, http.StatusBadRequest, response.Error("empty request body"))

	case errors.As(err, &syntaxErr):
		log.Warn("malformed JSON", slog.Int64("offset", syntaxErr.Offset))
		EncodeJSON(w, r, http.StatusBadRequest, response.Error("malformed JSON"))

	case errors.As(err, &unmarshalErr):
		log.Warn("invalid type for field",
			slog.String("field", unmarshalErr.Field),
			slog.String("value", unmarshalErr.Value),
		)
		msg := fmt.Sprintf("invalid type for field %s", unmarshalErr.Field)
		EncodeJSON(w, r, http.StatusBadRequest, response.Error(msg))

	default:
		log.Warn("invalid JSON", slog.Any("error", err))
		EncodeJSON(w, r, http.StatusBadRequest, response.Error("invalid JSON"))
	}

	return false
}

func Validate(w http.ResponseWriter, r *http.Request, log *slog.Logger, val *validator.Validate, v any) bool {
	if err := val.Struct(v); err != nil {
		ve := err.(validator.ValidationErrors)
		log.Warn("validation error", slog.Any("error", err))

		w.WriteHeader(http.StatusUnprocessableEntity)
		render.JSON(w, r, response.ValidationError(ve))

		return false
	}
	return true
}
