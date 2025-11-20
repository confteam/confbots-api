package handler

import (
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/confteam/confbots-api/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type BotHandler struct {
	uc  *usecase.BotUseCase
	log *slog.Logger
}

func NewBotHandler(uc *usecase.BotUseCase, log *slog.Logger) *BotHandler {
	return &BotHandler{
		uc:  uc,
		log: log,
	}
}

func (h *BotHandler) RegisterRoutes(r chi.Router) {
	r.Post("/bots", h.CreateIfNotExists)
}

const pkg = "transport.http.handler.BotHandler"

func (h *BotHandler) CreateIfNotExists(w http.ResponseWriter, r *http.Request) {
	const op = pkg + ".CreateIfNotExists"

	h.log = h.log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req dto.CreateBotRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.log.Error("failed to decode request body", slog.Any("error", err))
		render.JSON(w, r, resp.Error("failed to decode request"))
		return
	}

	h.log.Info("request body decoded",
		slog.Int("tgid", int(req.TgId)),
		slog.String("type", string(req.BotType)),
	)

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		h.log.Error("invalid request", slog.Any("error", err))
		render.JSON(w, r, resp.ValidationError(validateErr))
		return
	}

	bot, err := h.uc.CreateIfNotExists(r.Context(), req.TgId, req.BotType)
	if err != nil {
		h.log.Error("failed to create bot", slog.Any("error", err))
		render.JSON(w, r, resp.Error("failed to create bot"))
		return
	}

	response := dto.CreateBotResponse{
		Response:  resp.OK(),
		ID:        bot.ID,
		TgId:      bot.TgID,
		Type:      bot.Type,
		ChannelID: bot.ChannelID,
	}

	h.log.Info("bot authorized",
		slog.Int("id", int(response.ID)),
		slog.Int("tgid", int(response.TgId)),
		slog.String("type", string(response.Type)),
		slog.Any("channel_id", response.ChannelID),
	)

	render.JSON(w, r, response)
}
