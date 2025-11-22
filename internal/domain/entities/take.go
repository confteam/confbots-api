package entities

type Take struct {
	ID             int
	Status         string
	UserMessageID  int64
	AdminMessageID int64
	UserChannelID  int
	ChannelID      int
}
