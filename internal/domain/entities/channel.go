package entities

type Channel struct {
	ID                int32
	Code              string
	ChannelChatID     *int64
	AdminChatID       *int64
	DiscussionsChatID *int64
	Decorations       *string
}

type ChannelWithoutIDAndCode struct {
	ChannelChatID     *int64
	AdminChatID       *int64
	DiscussionsChatID *int64
	Decorations       *string
}
