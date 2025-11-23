package domain

import "context"

type Channel struct {
	ID                int
	Code              string
	ChannelChatID     *int64
	AdminChatID       *int64
	DiscussionsChatID *int64
	Decorations       *string
}

type ChannelWithoutCode struct {
	ID                int
	ChannelChatID     *int64
	AdminChatID       *int64
	DiscussionsChatID *int64
	Decorations       *string
}

type ChannelWithBotTgidAndType struct {
	BotTgID           int
	BotType           string
	Code              string
	ChannelChatID     *int64
	AdminChatID       *int64
	DiscussionsChatID *int64
	Decorations       *string
}

type ChannelRepository interface {
	Create(ctx context.Context, channel ChannelWithBotTgidAndType) (*Channel, error)
	Update(ctx context.Context, channel ChannelWithoutCode) (*Channel, error)
}
