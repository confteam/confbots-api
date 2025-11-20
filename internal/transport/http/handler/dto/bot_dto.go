package dto

import (
	"github.com/confteam/confbots-api/internal/domain/entities"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type CreateBotRequest struct {
	TgId    int32            `json:"tgid" validate:"required"`
	BotType entities.BotType `json:"type" validate:"required"`
}

type CreateBotResponse struct {
	ID        int32            `json:"id"`
	TgId      int32            `json:"tgid"`
	Type      entities.BotType `json:"type"`
	ChannelID *int32           `json:"channel_id,omitempty"`
	resp.Response
}
