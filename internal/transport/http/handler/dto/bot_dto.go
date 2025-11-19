package dto

import "github.com/confteam/confbots-api/internal/domain/entities"

type CreateBotRequest struct {
	TgId    int32            `json:"tgid"`
	BotType entities.BotType `json:"type"`
}

type CreateBotResponse struct {
	ID        int32            `json:"id"`
	TgId      int32            `json:"tgid"`
	Type      entities.BotType `json:"type"`
	ChannelID int32            `json:"channel_id"`
}
