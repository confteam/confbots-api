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

type UserHandler struct {
	uc  *usecase.UserUseCase
	log *slog.Logger
	val *validator.Validate
}

func NewUserHandler(uc *usecase.UserUseCase, log *slog.Logger) *UserHandler {
	return &UserHandler{
		uc:  uc,
		log: log,
		val: validator.New(),
	}
}

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Post("/users/{tgid}", h.Upsert)
	r.Patch("/users/role", h.UpdateRole)
	r.Get("/users/role", h.GetRole)
	r.Patch("/users/anonimity", h.ToggleAnonimity)
	r.Get("/users/anonimity", h.GetAnonimity)
	r.Get("/users", h.GetAllUsersInChannel)
}

const userPkg = "transport.http.handler.UserHandler"

func (h *UserHandler) GetQueries(w http.ResponseWriter, r *http.Request, log *slog.Logger) (int64, int, bool) {
	tgID, ok := helpers.ParseQuery(w, r, log, "tgId", true)
	channelID, ok := helpers.ParseQuery(w, r, log, "channelId", true)

	log.Info("query parameteres decoded",
		slog.Int("tgid", tgID),
		slog.Int("channel_id", channelID),
	)

	return int64(tgID), channelID, ok
}

func (h *UserHandler) Upsert(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".Upsert"

	log := helpers.ReqLogger(h.log, r, op)

	var req dto.UpsertUserRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
		return
	}

	tgID, ok := helpers.ParseURLParam(w, r, log, "tgid")
	if !ok {
		return
	}

	log.Info("request body decoded",
		slog.Int64("tgid", int64(tgID)),
		slog.Int("channel_id", req.ChannelID),
	)

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	userID, err := h.uc.Upsert(r.Context(), int64(tgID), req.ChannelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("upserted user",
		slog.Int64("tgid", int64(tgID)),
		slog.Int("user_id", userID),
		slog.Int("channel_id", req.ChannelID),
	)

	response := dto.UpsertUserResponse{
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".UpdateRole"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, channelID, ok := h.GetQueries(w, r, log)
	if !ok {
		return
	}

	var req dto.UpdateUserRoleRequest
	if !helpers.DecodeJSON(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.String("role", string(req.Role)),
	)

	if !helpers.Validate(w, r, log, h.val, req) {
		return
	}

	if err := h.uc.UpdateRole(r.Context(), req.Role, tgID, channelID); err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("updated user's role",
		slog.String("role", string(req.Role)),
		slog.Int64("tgid", tgID),
		slog.Int("channel_id", channelID),
	)

	response := dto.UpsertUserResponse{
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *UserHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".GetRole"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, channelID, ok := h.GetQueries(w, r, log)
	if !ok {
		return
	}

	role, err := h.uc.GetRole(r.Context(), tgID, channelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("got user's role",
		slog.String("role", string(role)),
		slog.Int64("tgid", tgID),
		slog.Int("channel_id", channelID),
	)

	response := dto.GetUserRoleResponse{
		Role:     role,
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *UserHandler) GetAnonimity(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".GetAnonimity"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, channelID, ok := h.GetQueries(w, r, log)
	if !ok {
		return
	}

	anonimity, err := h.uc.GetAnonimity(r.Context(), tgID, channelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("got user's anonimity",
		slog.Bool("anonimity", anonimity),
		slog.Int64("tgid", tgID),
		slog.Int("channel_id", channelID),
	)

	response := dto.UserAnonimityResponse{
		Anonimity: anonimity,
		Response:  resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *UserHandler) ToggleAnonimity(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".ToggleAnonimity"

	log := helpers.ReqLogger(h.log, r, op)

	tgID, channelID, ok := h.GetQueries(w, r, log)
	if !ok {
		return
	}

	anonimity, err := h.uc.ToggleAnonimity(r.Context(), tgID, channelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("toggled user's anonimity",
		slog.Bool("anonimity", anonimity),
		slog.Int64("tgid", tgID),
		slog.Int("channel_id", channelID),
	)

	response := dto.UserAnonimityResponse{
		Anonimity: anonimity,
		Response:  resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}

func (h *UserHandler) GetAllUsersInChannel(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".GetAllUsersInChannel"

	log := helpers.ReqLogger(h.log, r, op)

	channelID, ok := helpers.ParseQuery(w, r, log, "channelId", true)
	if !ok {
		return
	}

	tgIDs, err := h.uc.GetAllUsersInChannel(r.Context(), channelID)
	if err != nil {
		helpers.HandleError(w, r, log, err)
		return
	}

	log.Info("got all users in channel",
		slog.Int("channel_id", channelID),
		slog.Any("tgids", tgIDs),
	)

	response := dto.GetAllUsersInChannelResponse{
		TgIDs:    tgIDs,
		Response: resp.OK(),
	}

	helpers.EncodeJSON(w, r, http.StatusOK, response)
}
