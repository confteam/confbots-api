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

type ChannelWithoutID struct {
	Code              string
	ChannelChatID     *int64
	AdminChatID       *int64
	DiscussionsChatID *int64
	Decorations       *string
}

type ChannelRepository interface {
	Create(ctx context.Context, channel ChannelWithoutID) (int, error)
	Update(ctx context.Context, channel ChannelWithoutCode) (*Channel, error)
	FindByCode(ctx context.Context, code string) (*ChannelWithoutCode, error)
	FindByID(ctx context.Context, id int) (*ChannelWithoutID, error)
}
