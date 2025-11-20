package dto

import (
	"github.com/confteam/confbots-api/internal/domain/entities"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type CreateBotRequest struct {
	TgId    int64            `json:"tgid" validate:"required"`
	BotType entities.BotType `json:"type" validate:"required"`
}

type CreateBotResponse struct {
	ID      int32            `json:"id"`
	TgId    int64            `json:"tgid"`
	Type    entities.BotType `json:"type"`
	Channel *ChannelResponse `json:"channel,omitempty"`
	resp.Response
}

type ChannelResponse struct {
	ID                int32   `json:"id"`
	Code              string  `json:"code"`
	ChannelChatID     *int64  `json:"channel_chat_id"`
	AdminChatID       *int64  `json:"admin_chat_id"`
	DiscussionsChatID *int64  `json:"discussions_chat_id"`
	Decorations       *string `json:"decorations"`
}
