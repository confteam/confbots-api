package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	"github.com/confteam/confbots-api/internal/usecase"
	"github.com/go-chi/chi/v5"
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

func (h *BotHandler) CreateIfNotExists(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	bot, err := h.uc.CreateIfNotExists(r.Context(), req.TgId, req.BotType)
	if err != nil {
		h.log.Error("failed to create bot", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	resp := dto.CreateBotResponse{
		ID:        bot.ID,
		TgId:      bot.TgID,
		Type:      bot.Type,
		ChannelID: *bot.ChannelID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
