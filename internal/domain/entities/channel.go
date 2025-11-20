package entities

type Channel struct {
	ID                int32
	Code              string
	ChannelChatID     *int32
	AdminChatID       *int32
	DiscussionsChatID *int32
	Decorations       *string
}
