package domain

import "context"

type Take struct {
	ID             int
	Status         string
	UserMessageID  int64
	AdminMessageID int64
	UserChannelID  int
	ChannelID      int
}

type TakeRepository interface {
	Create(ctx context.Context, userMessageID int64, adminMessageID int64, userChannelID int, channelID int) (*Take, error)
	GetByID(ctx context.Context, id int) (*Take, error)
	GetByMsgID(ctx context.Context, messageID int64, channelID int) (*Take, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}
