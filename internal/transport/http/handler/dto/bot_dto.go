package dto

import (
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type CreateBotRequest struct {
	TgId    int64  `json:"tgid" validate:"required"`
	BotType string `json:"type" validate:"required"`
}

type CreateBotResponse struct {
	ID      int32            `json:"id"`
	TgId    int64            `json:"tgid"`
	Type    string           `json:"type"`
	Channel *ChannelResponse `json:"channel,omitempty"`
	resp.Response
}
