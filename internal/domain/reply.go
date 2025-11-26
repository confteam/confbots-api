package domain

import "context"

type Reply struct {
	ID             int
	UserMessageID  int64
	AdminMessageID int64
	TakeID         int
	ChannelID      int
}

type ReplyRepository interface {
	Create(ctx context.Context, userMessageID int64, adminMessageID int64, takeID int, channelID int) (int, error)
	GetByMsgIDAndChannelID(ctx context.Context, messageID int64, channelID int) (*Reply, error)
	GetByMsgIDAndTakeID(ctx context.Context, messageID int64, takeID int) (*Reply, error)
}
