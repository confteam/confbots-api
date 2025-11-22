package entities

type Reply struct {
	ID             int
	UserMessageID  int64
	AdminMessageID int64
	TakeID         int
	ChannelID      int
}
