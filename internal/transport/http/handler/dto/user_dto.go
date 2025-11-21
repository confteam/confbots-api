package dto

import (
	"github.com/confteam/confbots-api/internal/domain/entities"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type UpsertUserRequest struct {
	ChannelID int           `json:"channelId" validate:"required"`
	Role      entities.Role `json:"role" validate:"required"`
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

type UserAnonimityResponse struct {
	Anonimity bool `json:"anonimity"`
	resp.Response
}
