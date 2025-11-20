package dto

import resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"

// for bot
type ChannelResponse struct {
	ID                int32   `json:"id"`
	Code              string  `json:"code"`
	ChannelChatID     *int64  `json:"channel_chat_id"`
	AdminChatID       *int64  `json:"admin_chat_id"`
	DiscussionsChatID *int64  `json:"discussions_chat_id"`
	Decorations       *string `json:"decorations"`
}

type CreateChannelRequest struct {
	ChannelChatID     *int64 `json:"channel_chat_id,omitempty"`
	AdminChatID       *int64 `json:"admin_chat_id,omitempty"`
	DiscussionsChatID *int64 `json:"discussions_chat_id,omitempty"`
}

type CreateChannelResponse struct {
	resp.Response `json:"response"`
	ID            int32 `json:"id"`
}

type UpdateChannelRequest struct {
	ChannelChatID     *int64  `json:"channel_chat_id,omitempty"`
	AdminChatID       *int64  `json:"admin_chat_id,omitempty"`
	DiscussionsChatID *int64  `json:"discussions_chat_id,omitempty"`
	Decorations       *string `json:"decorations,omitempty"`
}

type UpdateChannelResponse struct {
	resp.Response
}
