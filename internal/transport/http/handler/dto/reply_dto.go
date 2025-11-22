package dto

import resp "github.com/confteam/confbots-api/internal/transport/http/handler/response"

type ReplyResponse struct {
	ID             int   `json:"id"`
	UserMessageID  int64 `json:"userMessageId"`
	AdminMessageID int64 `json:"adminMessageId"`
	TakeID         int   `json:"takeId"`
	ChannelID      int   `json:"channelId"`
}

type CreateReplyRequest struct {
	UserMessageID  int64 `json:"userMessageId" validate:"required"`
	AdminMessageID int64 `json:"adminMessageId" validate:"required"`
	TakeID         int   `json:"takeId" validate:"required"`
	ChannelID      int   `json:"channelId" validate:"required"`
}

type CreateReplyResponse struct {
	ID int `json:"id"`
	resp.Response
}

type GetReplyByMsgIDResponse struct {
	Reply ReplyResponse `json:"reply"`
	resp.Response
}
