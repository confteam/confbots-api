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
	r.Get("/replies/{tgid}", h.GetRouterQuery)
}

const replyPkg = "transport.http.handler.ReplyHandler"

func (h *ReplyHandler) GetRouterQuery(w http.ResponseWriter, r *http.Request) {
	const op = replyPkg + ".GetRouterQuery"

	log := helpers.ReqLogger(h.log, r, op)

	channelID, _ := helpers.ParseQuery(w, r, log, "channelId", false)
	takeID, _ := helpers.ParseQuery(w, r, log, "takeId", false)

	if channelID != 0 {
		h.GetByMsgIDAndChannelID(w, r, channelID)
	} else {
		h.GetByMsgIDAndTakeID(w, r, takeID)
	}
}

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

func (h *ReplyHandler) GetByMsgIDAndChannelID(w http.ResponseWriter, r *http.Request, channelID int) {
	const op = replyPkg + ".GetByMsgID"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, ok := helpers.ParseURLParam(w, r, log, "tgid")
	if !ok {
		return
	}

	log.Info("got url params", slog.Int64("tgid", int64(tgID)), slog.Int("channel_id", channelID))

	reply, err := h.uc.GetByMsgIDAndChannelID(r.Context(), int64(tgID), channelID)
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

func (h *ReplyHandler) GetByMsgIDAndTakeID(w http.ResponseWriter, r *http.Request, takeID int) {
	const op = replyPkg + ".GetByMsgID"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, ok := helpers.ParseURLParam(w, r, log, "tgid")
	if !ok {
		return
	}

	log.Info("got url params", slog.Int64("tgid", int64(tgID)), slog.Int("take_id", takeID))

	reply, err := h.uc.GetByMsgIDAndTakeID(r.Context(), int64(tgID), takeID)
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
