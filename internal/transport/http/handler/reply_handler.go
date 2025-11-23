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

type ReplyHandler struct {
	uc  *usecase.ReplyUseCase
	log *slog.Logger
	val *validator.Validate
}

func NewReplyHandler(uc *usecase.ReplyUseCase, log *slog.Logger) *ReplyHandler {
	return &ReplyHandler{
		uc:  uc,
		log: log,
		val: validator.New(),
	}
}

func (h *ReplyHandler) RegisterRoutes(r chi.Router) {
	r.Post("/replies", h.Create)
	r.Get("/replies/{tgid}", h.GetByMsgID)
}

const replyPkg = "transport.http.handler.ReplyHandler"

func (h *ReplyHandler) Create(w http.ResponseWriter, r *http.Request) {
	const op = replyPkg + ".Create"

	log := helpers.ReqLogger(h.log, r, op)

	var req dto.CreateReplyRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.Int64("user_message_id", req.UserMessageID),
		slog.Int64("admin_message_id", req.AdminMessageID),
		slog.Int("take_id", req.TakeID),
		slog.Int("channel_id", req.ChannelID),
	)

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	id, err := h.uc.Create(r.Context(), req.UserMessageID, req.AdminMessageID, req.TakeID, req.ChannelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("created reply",
		slog.Int("id", id),
		slog.Int64("user_message_id", req.UserMessageID),
		slog.Int64("admin_message_id", req.AdminMessageID),
		slog.Int("take_id", req.TakeID),
		slog.Int("channel_id", req.ChannelID),
	)

	response := dto.CreateReplyResponse{
		ID:       id,
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *ReplyHandler) GetByMsgID(w http.ResponseWriter, r *http.Request) {
	const op = replyPkg + ".GetByMsgID"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, ok := helpers.ParseURLParam(w, r, log, "tgid")
	if !ok {
		return
	}

	channelID, ok := helpers.ParseQuery(w, r, log, "channelId", true)
	if !ok {
		return
	}

	log.Info("got url params", slog.Int("tgid", tgID), slog.Int("channel_id", channelID))

	reply, err := h.uc.GetByMsgID(r.Context(), int64(tgID), channelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("got reply",
		slog.Int("id", reply.ID),
		slog.Int64("user_message_id", reply.UserMessageID),
		slog.Int64("admin_message_id", reply.AdminMessageID),
		slog.Int("take_id", reply.TakeID),
		slog.Int("channel_id", reply.ChannelID),
	)

	response := dto.GetReplyByMsgIDResponse{
		Reply:    helpers.MapReplyToReplyResponse(*reply),
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}
