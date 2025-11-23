package handler

import (
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/confteam/confbots-api/internal/transport/http/helpers"
	"github.com/confteam/confbots-api/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type BotHandler struct {
	uc  *usecase.BotUseCase
	log *slog.Logger
	val *validator.Validate
}

func NewBotHandler(uc *usecase.BotUseCase, log *slog.Logger) *BotHandler {
	val := validator.New()
	return &BotHandler{
		uc:  uc,
		log: log,
		val: val,
	}
}

func (h *BotHandler) RegisterRoutes(r chi.Router) {
	r.Post("/bots", h.Auth)
}

const botPkg = "transport.http.handler.BotHandler"

func (h *BotHandler) Auth(w http.ResponseWriter, r *http.Request) {
	const op = botPkg + ".Auth"

	log := helpers.ReqLogger(h.log, r, op)

	var req dto.CreateBotRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.Int64("tgid", req.TgId),
		slog.String("type", string(req.BotType)),
	)

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	bot, err := h.uc.Auth(r.Context(), req.TgId, req.BotType)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	var channel *dto.ChannelResponse
	if bot.Channel != nil {
		channel = &dto.ChannelResponse{
			ID:                bot.Channel.ID,
			Code:              bot.Channel.Code,
			ChannelChatID:     bot.Channel.ChannelChatID,
			AdminChatID:       bot.Channel.AdminChatID,
			DiscussionsChatID: bot.Channel.DiscussionsChatID,
			Decorations:       bot.Channel.Decorations,
		}
	}

	response := dto.CreateBotResponse{
		Response: resp.OK(),
		ID:       bot.ID,
		TgId:     bot.TgID,
		Type:     bot.Type,
		Channel:  channel,
	}

	log.Info("bot authorized",
		slog.Int("id", int(response.ID)),
		slog.Int("tgid", int(response.TgId)),
		slog.String("type", string(response.Type)),
		slog.Any("channel", channel),
	)

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}
