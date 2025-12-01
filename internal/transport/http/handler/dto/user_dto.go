package dto

import (
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type UpsertUserRequest struct {
	ChannelID int `json:"channelId" validate:"required"`
}

type UpsertUserResponse struct {
	resp.Response
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type UpdateUserRoleResponse struct {
	resp.Response
}

type GetUserRoleResponse struct {
	Role string `json:"role"`
	resp.Response
}

type UserAnonimityResponse struct {
	Anonimity bool `json:"anonimity"`
	resp.Response
}

type GetAllUsersInChannelResponse struct {
	TgIDs []int64 `json:"tgids"`
	resp.Response
}
