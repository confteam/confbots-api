package handler

import (
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func reqLogger(base *slog.Logger, r *http.Request, op string) *slog.Logger {
	return base.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
}

func json(w http.ResponseWriter, r *http.Request, status int, v any) {
	w.WriteHeader(status)
	render.JSON(w, r, v)
}

func returnError(w http.ResponseWriter, r *http.Request, log *slog.Logger, status int, msg string, err error) {
	if err != nil {
		log.Error(msg, slog.Any("error", err))
	}
	json(w, r, status, resp.Error(msg))
}

func decodeJson[T any](w http.ResponseWriter, r *http.Request, log *slog.Logger, dst *T) bool {
	if err := render.DecodeJSON(r.Body, dst); err != nil {
		returnError(w, r, log, http.StatusBadRequest, "failed to decode request", err)
		return false
	}
	return true
}

func validate(w http.ResponseWriter, r *http.Request, log *slog.Logger, val *validator.Validate, v any) bool {
	if err := val.Struct(v); err != nil {
		ve := err.(validator.ValidationErrors)
		log.Error("validation error", slog.Any("error", err))
		w.WriteHeader(http.StatusUnprocessableEntity)
		render.JSON(w, r, resp.ValidationError(ve))
		return false
	}
	return true
}

func mapChannelToChannelResponse(channel entities.Channel) dto.ChannelResponse {
	return dto.ChannelResponse{
		ID:                channel.ID,
		Code:              channel.Code,
		ChannelChatID:     channel.ChannelChatID,
		AdminChatID:       channel.AdminChatID,
		DiscussionsChatID: channel.DiscussionsChatID,
		Decorations:       channel.Decorations,
	}
}

func mapTakeToTakeResponse(take entities.Take) dto.TakeResponse {
	return dto.TakeResponse{
		ID:             take.ID,
		Status:         take.Status,
		UserMessageID:  take.UserMessageID,
		AdminMessageID: take.AdminMessageID,
		UserChannelID:  take.UserChannelID,
		ChannelID:      take.ChannelID,
	}
}

func mapReplyToReplyResponse(reply entities.Reply) dto.ReplyResponse {
	return dto.ReplyResponse{
		ID:             reply.ID,
		UserMessageID:  reply.UserMessageID,
		AdminMessageID: reply.AdminMessageID,
		TakeID:         reply.TakeID,
		ChannelID:      reply.ChannelID,
	}
}
