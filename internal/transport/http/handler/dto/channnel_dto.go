package dto

import (
	resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"
)

type ChannelResponse struct {
	ID                int     `json:"id"`
	Code              string  `json:"code"`
	ChannelChatID     *int64  `json:"channelChatId"`
	AdminChatID       *int64  `json:"adminChatId"`
	DiscussionsChatID *int64  `json:"discussionsChatId"`
	Decorations       *string `json:"decorations"`
}

type ChannelWithoutCodeResponse struct {
	ID                int     `json:"id"`
	ChannelChatID     *int64  `json:"channelChatId"`
	AdminChatID       *int64  `json:"adminChatId"`
	DiscussionsChatID *int64  `json:"discussionsChatId"`
	Decorations       *string `json:"decorations"`
}

type ChannelWithoutIDResponse struct {
	Code              string  `json:"code"`
	ChannelChatID     *int64  `json:"channelChatId"`
	AdminChatID       *int64  `json:"adminChatId"`
	DiscussionsChatID *int64  `json:"discussionsChatId"`
	Decorations       *string `json:"decorations"`
}

type CreateChannelRequest struct {
	Code              string `json:"code" validate:"required"`
	ChannelChatID     *int64 `json:"channelChatId,omitempty"`
	AdminChatID       *int64 `json:"adminChatId,omitempty"`
	DiscussionsChatID *int64 `json:"discussionsChatId,omitempty"`
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

type FindChannelByCodeResponse struct {
	resp.Response
	Channel ChannelWithoutCodeResponse `json:"channel"`
}

type FindChannelByIDResponse struct {
	resp.Response
	Channel ChannelWithoutIDResponse `json:"channel"`
}
