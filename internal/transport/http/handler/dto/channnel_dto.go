package dto

import (
	"github.com/confteam/confbots-api/internal/domain/entities"
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

// for bot
type ChannelResponse struct {
	ID                int     `json:"id"`
	Code              string  `json:"code"`
	ChannelChatID     *int64  `json:"channelChatId"`
	AdminChatID       *int64  `json:"adminChatId"`
	DiscussionsChatID *int64  `json:"discussionsChatId"`
	Decorations       *string `json:"decorations"`
}

type CreateChannelRequest struct {
	BotTgId           int64            `json:"botTgId" validate:"required"`
	BotType           entities.BotType `json:"botType" validate:"required"`
	Code              string           `json:"code" validate:"required"`
	ChannelChatID     *int64           `json:"channelChatId,omitempty"`
	AdminChatID       *int64           `json:"adminChatId,omitempty"`
	DiscussionsChatID *int64           `json:"discussionsChatId,omitempty"`
}

type CreateChannelResponse struct {
	resp.Response `json:"response"`
	ID            int `json:"id"`
}

type UpdateChannelRequest struct {
	ChannelChatID     *int64  `json:"channelChatId,omitempty"`
	AdminChatID       *int64  `json:"adminChatId,omitempty"`
	DiscussionsChatID *int64  `json:"discussionsChatId,omitempty"`
	Decorations       *string `json:"decorations,omitempty"`
}

type UpdateChannelResponse struct {
	resp.Response
}
