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
	r.Post("/users/{id}", h.Upsert)
	r.Patch("/users/role", h.UpdateRole)
	r.Get("/users/role", h.GetRole)
	r.Patch("/users/anonimity", h.ToggleAnonimity)
	r.Get("/users/anonimity", h.GetAnonimity)
}

const userPkg = "transport.http.handler.UserHandler"

func (h *UserHandler) GetQueries(w http.ResponseWriter, r *http.Request, log *slog.Logger) (int64, int) {
	tgIDStr := r.URL.Query().Get("tgId")
	channelIDStr := r.URL.Query().Get("channelId")
	if tgIDStr == "" || channelIDStr == "" {
		returnError(w, r, log, http.StatusBadRequest, "tgId and channelId are required", nil)
		return 0, 0
	}

	tgID, err := strconv.Atoi(tgIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert tgid", err)
		return 0, 0
	}

	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert channelId", err)
		return 0, 0
	}

	log.Info("query parameteres decoded",
		slog.Int("tgid", tgID),
		slog.Int("channel_id", channelID),
	)

	return int64(tgID), channelID
}

func (h *UserHandler) Upsert(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".Upsert"

	log := reqLogger(h.log, r, op)

	var req dto.UpsertUserRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	tgIDStr := chi.URLParam(r, "id")
	tgID, err := strconv.Atoi(tgIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert tgID", nil)
		return
	}

	log.Info("request body decoded",
		slog.String("role", string(req.Role)),
		slog.Int64("tgid", int64(tgID)),
		slog.Int("channel_id", req.ChannelID),
	)

	if !validate(w, r, log, h.val, req) {
		return
	}

	userID, err := h.uc.Upsert(r.Context(), int64(tgID), req.ChannelID, req.Role)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to upsert user", err)
		return
	}

	log.Info("upserted user",
		slog.Int64("tgid", int64(tgID)),
		slog.Int("user_id", userID),
		slog.Int("channel_id", req.ChannelID),
		slog.String("role", string(req.Role)),
	)

	response := dto.UpsertUserResponse{
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".UpdateRole"

	log := reqLogger(h.log, r, op)

	tgID, channelID := h.GetQueries(w, r, log)
	if tgID == 0 || channelID == 0 {
		return
	}

	var req dto.UpdateUserRoleRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.String("role", string(req.Role)),
	)

	if !validate(w, r, log, h.val, req) {
		return
	}

	if err := h.uc.UpdateRole(r.Context(), req.Role, tgID, channelID); err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to update user's role", err)
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

	json(w, r, http.StatusOK, response)
}

func (h *UserHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".GetRole"

	log := reqLogger(h.log, r, op)

	tgID, channelID := h.GetQueries(w, r, log)
	if tgID == 0 || channelID == 0 {
		return
	}

	role, err := h.uc.GetRole(r.Context(), tgID, channelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to get user's role", err)
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

	json(w, r, http.StatusOK, response)
}

func (h *UserHandler) GetAnonimity(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".GetAnonimity"

	log := reqLogger(h.log, r, op)

	tgID, channelID := h.GetQueries(w, r, log)
	if tgID == 0 || channelID == 0 {
		return
	}

	anonimity, err := h.uc.GetAnonimity(r.Context(), tgID, channelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to get user's anonimity", err)
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

	json(w, r, http.StatusOK, response)
}

func (h *UserHandler) ToggleAnonimity(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".ToggleAnonimity"

	log := reqLogger(h.log, r, op)

	tgID, channelID := h.GetQueries(w, r, log)
	if tgID == 0 || channelID == 0 {
		return
	}

	anonimity, err := h.uc.ToggleAnonimity(r.Context(), tgID, channelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to toggle user's anonimity", err)
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

	json(w, r, http.StatusOK, response)
}
