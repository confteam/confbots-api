package dto

import (
	"github.com/confteam/confbots-api/internal/domain/entities"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type UpsertUserRequest struct {
	TgID      int64 `json:"tgid" validate:"required"`
	ChannelID int   `json:"channel_id" validate:"required"`
}

type UpsertUserResponse struct {
	resp.Response
}

type UpdateUserRoleRequest struct {
	Role entities.Role `json:"role" validate:"required"`
}

type UpdateUserRoleResponse struct {
	resp.Response
}

type GetUserRoleResponse struct {
	Role entities.Role `json:"role"`
	resp.Response
}
