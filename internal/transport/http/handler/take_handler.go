package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
	"github.com/confteam/confbots-api/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type TakeHandler struct {
	uc  *usecase.TakeUseCase
	log *slog.Logger
	val *validator.Validate
}

func NewTakeHandler(uc *usecase.TakeUseCase, log *slog.Logger) *TakeHandler {
	return &TakeHandler{
		uc:  uc,
		log: log,
		val: validator.New(),
	}
}

func (h *TakeHandler) RegisterRoutes(r chi.Router) {
	r.Post("/takes", h.Create)
	r.Get("/takes/{id}", h.GetById)
	r.Get("/takes", h.GetByMsgId)
	r.Patch("/takes/{id}/status", h.UpdateStatus)
	r.Get("/takes/{id}/author", h.GetAuthor)
}

const takePkg = "transport.http.handler.TakeHandler"

func (h *TakeHandler) GetIDAndChannelID(w http.ResponseWriter, r *http.Request, log *slog.Logger) (int, int) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert id", nil)
		return 0, 0
	}

	channelIDStr := r.URL.Query().Get("channelId")
	if channelIDStr == "" {
		returnError(w, r, log, http.StatusBadRequest, "channelId is required", nil)
		return 0, 0
	}

	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert channelID", nil)
		return 0, 0
	}

	log.Info("got url params",
		slog.Int("id", id),
		slog.Int("channel_id", channelID),
	)

	return id, channelID
}

func (h *TakeHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".Create"

	log := reqLogger(h.log, r, op)

	var req dto.CreateTakeRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.Int64("user_tg_id", req.UserTgID),
		slog.Int64("user_message_id", req.UserMessageID),
		slog.Int64("admin_message_id", req.AdminMessageID),
		slog.Int("channel_id", req.ChannelID),
	)

	if !validate(w, r, log, h.val, req) {
		return
	}

	take, err := h.uc.Create(r.Context(), req.UserTgID, req.UserMessageID, req.AdminMessageID, req.ChannelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to create take", err)
		return
	}

	log.Info("created take",
		slog.Int("id", take.ID),
		slog.Int64("user_message_id", take.UserMessageID),
		slog.Int64("admin_message_id", take.AdminMessageID),
		slog.Int("user_channel_id", take.UserChannelID),
		slog.Int("channel_id", take.ChannelID),
	)

	response := dto.CreateTakeResponse{
		ID:       take.ID,
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}

func (h *TakeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".GetById"

	log := reqLogger(h.log, r, op)

	id, channelID := h.GetIDAndChannelID(w, r, log)
	if id == 0 || channelID == 0 {
		return
	}

	take, err := h.uc.GetById(r.Context(), id, channelID)
	if err != nil {
		returnError(w, r, log, http.StatusNotFound, "take not found", err)
		return
	}

	log.Info("got take",
		slog.Int("id", take.ID),
		slog.String("status", take.Status),
		slog.Int64("user_message_id", take.UserMessageID),
		slog.Int64("admin_message_id", take.AdminMessageID),
		slog.Int("user_channel_id", take.UserChannelID),
		slog.Int("channel_id", take.ChannelID),
	)

	response := dto.GetTakeResponse{
		Take:     mapTakeToTakeResponse(*take),
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}

func (h *TakeHandler) GetByMsgId(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".GetByMsgId"

	log := reqLogger(h.log, r, op)

	channelIDStr := r.URL.Query().Get("channelId")
	if channelIDStr == "" {
		returnError(w, r, log, http.StatusBadRequest, "channelId is required", nil)
		return
	}

	messageIDStr := r.URL.Query().Get("messageId")
	if messageIDStr == "" {
		returnError(w, r, log, http.StatusBadRequest, "messageId is required", nil)
		return
	}

	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert channelID", nil)
		return
	}

	messageID, err := strconv.Atoi(messageIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert messageID", nil)
		return
	}

	log.Info("got url params",
		slog.Int("message_id", messageID),
		slog.Int("channel_id", channelID),
	)

	take, err := h.uc.GetByMsgId(r.Context(), int64(messageID), channelID)
	if err != nil {
		returnError(w, r, log, http.StatusNotFound, "take not found", err)
		return
	}

	log.Info("got take",
		slog.Int("id", take.ID),
		slog.String("status", take.Status),
		slog.Int64("user_message_id", take.UserMessageID),
		slog.Int64("admin_message_id", take.AdminMessageID),
		slog.Int("user_channel_id", take.UserChannelID),
		slog.Int("channel_id", take.ChannelID),
	)

	response := dto.GetTakeResponse{
		Take:     mapTakeToTakeResponse(*take),
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}

func (h *TakeHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".UpdateStatus"

	log := reqLogger(h.log, r, op)

	id, channelID := h.GetIDAndChannelID(w, r, log)
	if id == 0 || channelID == 0 {
		return
	}

	var req dto.UpdateTakeStatusRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	log.Info("request body decoded", slog.String("status", req.Status))

	if !validate(w, r, log, h.val, req) {
		return
	}

	if err := h.uc.UpdateStatus(r.Context(), id, channelID); err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to update take's status", nil)
		return
	}

	log.Info("updated take's status",
		slog.Int("id", id),
		slog.Int("channel_id", channelID),
		slog.String("status", req.Status),
	)

	json(w, r, http.StatusOK, resp.OK())
}

func (h *TakeHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".GetAuthor"

	log := reqLogger(h.log, r, op)

	id, channelID := h.GetIDAndChannelID(w, r, log)
	if id == 0 || channelID == 0 {
		return
	}

	tgid, err := h.uc.GetTakeAuthor(r.Context(), id, channelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to get take's author", nil)
		return
	}

	log.Info("updated take's status",
		slog.Int("id", id),
		slog.Int("channel_id", channelID),
		slog.Int64("tgid", tgid),
	)

	response := dto.GetTakeAuthorResponse{
		TgID:     tgid,
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}
