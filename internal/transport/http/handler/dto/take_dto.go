package dto

import resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"

type TakeResponse struct {
	ID             int    `json:"id"`
	Status         string `json:"status"`
	UserMessageID  int64  `json:"user_message_id"`
	AdminMessageID int64  `json:"admin_message_id"`
	UserChannelID  int    `json:"user_channel_id"`
	ChannelID      int    `json:"channel_id"`
}

type CreateTakeRequest struct {
	UserTgID       int64 `json:"userTgId" validate:"required"`
	UserMessageID  int64 `json:"userMessageId" validate:"required"`
	AdminMessageID int64 `json:"adminMessageId" validate:"required"`
	ChannelID      int   `json:"channelId" validate:"required"`
}

type CreateTakeResponse struct {
	resp.Response
	ID int `json:"id"`
}

type GetTakeResponse struct {
	Take TakeResponse `json:"take"`
	resp.Response
}
