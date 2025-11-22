package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/confteam/confbots-api/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ChannelHandler struct {
	uc  *usecase.ChannelUseCase
	log *slog.Logger
	val *validator.Validate
}

func NewChannelHandler(uc *usecase.ChannelUseCase, log *slog.Logger) *ChannelHandler {
	return &ChannelHandler{
		uc:  uc,
		log: log,
		val: validator.New(),
	}
}

func (h *ChannelHandler) RegisterRoutes(r chi.Router) {
	r.Post("/channels", h.Create)
	r.Patch("/channels/{id}", h.Update)
}

const channelPkg = "transport.http.handler.ChannelHandler"

func (h *ChannelHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = channelPkg + ".Create"

	log := reqLogger(h.log, r, op)

	var req dto.CreateChannelRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.Int64("bot_tgid", req.BotTgId),
		slog.String("bot_type", string(req.BotType)),
		slog.String("code", req.Code),
		slog.Int64("channel_chat_id", *req.ChannelChatID),
		slog.Int64("admin_chat_id", *req.AdminChatID),
		slog.Int64("discussions_chat_id", *req.DiscussionsChatID),
	)

	if !validate(w, r, log, h.val, req) {
		return
	}

	channel, err := h.uc.Create(r.Context(), entities.ChannelWithBotTgidAndType{
		BotType:           string(req.BotType),
		BotTgID:           int(req.BotTgId),
		Code:              req.Code,
		ChannelChatID:     req.ChannelChatID,
		AdminChatID:       req.AdminChatID,
		DiscussionsChatID: req.DiscussionsChatID,
	})
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to create channel", err)
		return
	}

	response := dto.CreateChannelResponse{
		Response: resp.OK(),
		Channel:  mapChannelToChannelResponse(*channel),
	}

	log.Info("channel created",
		slog.Int("id", int(channel.ID)),
		slog.String("code", channel.Code),
		slog.Int64("channel_chat_id", *channel.ChannelChatID),
		slog.Int64("admin_chat_id", *channel.AdminChatID),
		slog.Int64("discussions_chat_id", *channel.DiscussionsChatID),
		slog.String("decorations", *channel.Decorations),
	)

	json(w, r, http.StatusOK, response)
}

func (h *ChannelHandler) Update(w http.ResponseWriter, r *http.Request) {
	const op = channelPkg + ".Update"

	log := reqLogger(h.log, r, op)

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert id", nil)
		return
	}

	var req dto.UpdateChannelRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	if req.Decorations == nil {
		req.Decorations = new(string)
	}

	log.Info("request body decoded",
		slog.Int("id", id),
		slog.Int64("channel_chat_id", *req.ChannelChatID),
		slog.Int64("admin_chat_id", *req.AdminChatID),
		slog.Int64("discussions_chat_id", *req.DiscussionsChatID),
		slog.String("decorations", *req.Decorations),
	)

	if !validate(w, r, log, h.val, req) {
		return
	}

	channel, err := h.uc.Update(r.Context(), entities.ChannelWithoutCode{
		ID:                id,
		ChannelChatID:     req.ChannelChatID,
		AdminChatID:       req.AdminChatID,
		DiscussionsChatID: req.DiscussionsChatID,
		Decorations:       req.Decorations,
	})
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to update channel", err)
		return
	}

	log.Info("channel updated",
		slog.Int("id", int(channel.ID)),
		slog.String("code", channel.Code),
		slog.Int64("channel_chat_id", *channel.ChannelChatID),
		slog.Int64("admin_chat_id", *channel.AdminChatID),
		slog.Int64("discussions_chat_id", *channel.DiscussionsChatID),
		slog.String("decorations", *channel.Decorations),
	)

	response := dto.UpdateChannelResponse{
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}
