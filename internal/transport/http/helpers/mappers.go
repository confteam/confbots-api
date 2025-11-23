package helpers

import (
	"github.com/confteam/confbots-api/internal/domain"
	"github.com/confteam/confbots-api/internal/transport/http/handler/dto"
)

func MapChannelToChannelResponse(channel domain.Channel) dto.ChannelResponse {
	return dto.ChannelResponse{
		ID:                channel.ID,
		Code:              channel.Code,
		ChannelChatID:     channel.ChannelChatID,
		AdminChatID:       channel.AdminChatID,
		DiscussionsChatID: channel.DiscussionsChatID,
		Decorations:       channel.Decorations,
	}
}

func MapTakeToTakeResponse(take domain.Take) dto.TakeResponse {
	return dto.TakeResponse{
		ID:             take.ID,
		Status:         take.Status,
		UserMessageID:  take.UserMessageID,
		AdminMessageID: take.AdminMessageID,
		UserChannelID:  take.UserChannelID,
		ChannelID:      take.ChannelID,
	}
}

func MapReplyToReplyResponse(reply domain.Reply) dto.ReplyResponse {
	return dto.ReplyResponse{
		ID:             reply.ID,
		UserMessageID:  reply.UserMessageID,
		AdminMessageID: reply.AdminMessageID,
		TakeID:         reply.TakeID,
		ChannelID:      reply.ChannelID,
	}
}
