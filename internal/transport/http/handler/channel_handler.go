package handler

import (
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/domain"
	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/confteam/confbots-api/internal/transport/http/helpers"
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

	log := helpers.ReqLogger(h.log, r, op)

	var req dto.CreateChannelRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
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

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	channel, err := h.uc.Create(r.Context(), domain.ChannelWithBotTgidAndType{
		BotType:           string(req.BotType),
		BotTgID:           int(req.BotTgId),
		Code:              req.Code,
		ChannelChatID:     req.ChannelChatID,
		AdminChatID:       req.AdminChatID,
		DiscussionsChatID: req.DiscussionsChatID,
	})
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	response := dto.CreateChannelResponse{
		Response: resp.OK(),
		Channel:  helpers.MapChannelToChannelResponse(*channel),
	}

	log.Info("channel created",
		slog.Int("id", int(channel.ID)),
		slog.String("code", channel.Code),
		slog.Int64("channel_chat_id", *channel.ChannelChatID),
		slog.Int64("admin_chat_id", *channel.AdminChatID),
		slog.Int64("discussions_chat_id", *channel.DiscussionsChatID),
		slog.String("decorations", *channel.Decorations),
	)

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *ChannelHandler) Update(w http.ResponseWriter, r *http.Request) {
	const op = channelPkg + ".Update"

	log := helpers.ReqLogger(h.log, r, op)

	id, ok := helpers.ParseURLParam(w, r, log, "id")
	if !ok {
		return
	}

	var req dto.UpdateChannelRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
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

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	channel, err := h.uc.Update(r.Context(), domain.ChannelWithoutCode{
		ID:                id,
		ChannelChatID:     req.ChannelChatID,
		AdminChatID:       req.AdminChatID,
		DiscussionsChatID: req.DiscussionsChatID,
		Decorations:       req.Decorations,
	})
	if err != nil {
		helpers.HandleError(w, r, log, err)
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

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}
