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
	r.Post("/users", h.Upsert)
	r.Patch("/users/role", h.UpdateRole)
	r.Get("/users/role", h.GetRole)
}

const userPkg = "transport.http.handler.UserHandler"

func (h *UserHandler) GetQueries(w http.ResponseWriter, r *http.Request, log *slog.Logger) (int, int) {
	userIDStr := r.URL.Query().Get("userId")
	channelIDStr := r.URL.Query().Get("channelId")
	if userIDStr == "" || channelIDStr == "" {
		returnError(w, r, log, http.StatusBadRequest, "userId and channelId are required", nil)
		return 0, 0
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert userId", err)
		return 0, 0
	}

	channelID, err := strconv.Atoi(channelIDStr)
	if err != nil {
		returnError(w, r, log, http.StatusUnprocessableEntity, "failed to convert channelId", err)
		return 0, 0
	}

	log.Info("query parameteres decoded",
		slog.Int("user_id", userID),
		slog.Int("channel_id", channelID),
	)

	return userID, channelID
}

func (h *UserHandler) Upsert(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".Upsert"

	log := reqLogger(h.log, r, op)

	var req dto.UpsertUserRequest
	if !decodeJson(w, r, log, &req) {
		return
	}

	log.Info("request body decoded",
		slog.Int64("tgid", req.TgID),
		slog.Int("channel_id", req.ChannelID),
	)

	if !validate(w, r, log, h.val, req) {
		return
	}

	userID, err := h.uc.Upsert(r.Context(), req.TgID, req.ChannelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to upsert user", err)
		return
	}

	log.Info("upserted user",
		slog.Int64("tgid", req.TgID),
		slog.Int("user_id", userID),
		slog.Int("channel_id", req.ChannelID),
	)

	response := dto.UpsertUserResponse{
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}

func (h *UserHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	const op = userPkg + ".UpdateRole"

	log := reqLogger(h.log, r, op)

	userID, channelID := h.GetQueries(w, r, log)

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

	if err := h.uc.UpdateRole(r.Context(), req.Role, userID, channelID); err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to update user's role", err)
		return
	}

	log.Info("updated user's role",
		slog.String("role", string(req.Role)),
		slog.Int("user_id", userID),
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

	userID, channelID := h.GetQueries(w, r, log)

	role, err := h.uc.GetRole(r.Context(), userID, channelID)
	if err != nil {
		returnError(w, r, log, http.StatusInternalServerError, "failed to get user's role", err)
		return
	}

	log.Info("got user's role",
		slog.String("role", string(role)),
		slog.Int("user_id", userID),
		slog.Int("channel_id", channelID),
	)

	response := dto.GetUserRoleResponse{
		Role:     role,
		Response: resp.OK(),
	}

	json(w, r, http.StatusOK, response)
}
