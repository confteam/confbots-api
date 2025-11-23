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

func (h *TakeHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".Create"

	log := helpers.ReqLogger(h.log, r, op)

	var req dto.CreateTakeRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.Int64("user_tg_id", req.UserTgID),
		slog.Int64("user_message_id", req.UserMessageID),
		slog.Int64("admin_message_id", req.AdminMessageID),
		slog.Int("channel_id", req.ChannelID),
	)

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	take, err := h.uc.Create(r.Context(), req.UserTgID, req.UserMessageID, req.AdminMessageID, req.ChannelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
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

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *TakeHandler) GetById(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".GetById"

	log := helpers.ReqLogger(h.log, r, op)

	id, ok := helpers.ParseURLParam(w, r, log, "id")
	if !ok {
		return
	}

	take, err := h.uc.GetById(r.Context(), id)
	if err != nil {
		helpers.HandleError(w, r, log, err)
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
		Take:     helpers.MapTakeToTakeResponse(*take),
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *TakeHandler) GetByMsgId(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".GetByMsgId"

	log := helpers.ReqLogger(h.log, r, op)

	channelID, ok := helpers.ParseQuery(w, r, log, "channelId", true)
	messageID, ok := helpers.ParseQuery(w, r, log, "messageId", true)

	if !ok {
		return
	}

	log.Info("got url params",
		slog.Int("message_id", messageID),
		slog.Int("channel_id", channelID),
	)

	take, err := h.uc.GetByMsgId(r.Context(), int64(messageID), channelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
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
		Take:     helpers.MapTakeToTakeResponse(*take),
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *TakeHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".UpdateStatus"

	log := helpers.ReqLogger(h.log, r, op)

	id, ok := helpers.ParseURLParam(w, r, log, "id")
	if !ok {
		return
	}

	var req dto.UpdateTakeStatusRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
		return
	}

	log.Info("request body decoded", slog.String("status", req.Status))

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	if err := h.uc.UpdateStatus(r.Context(), id, req.Status); err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("updated take's status",
		slog.Int("id", id),
		slog.String("status", req.Status),
	)

	helpers.EncodeJSON(w, r, http.StatusOK, resp.OK())
}

func (h *TakeHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	const op = takePkg + ".GetAuthor"

	log := helpers.ReqLogger(h.log, r, op)

	id, ok := helpers.ParseURLParam(w, r, log, "id")
	if !ok {
		return
	}

	tgid, err := h.uc.GetTakeAuthor(r.Context(), id)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("got take's authour",
		slog.Int("id", id),
		slog.Int64("tgid", tgid),
	)

	response := dto.GetTakeAuthorResponse{
		TgID:     tgid,
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}
