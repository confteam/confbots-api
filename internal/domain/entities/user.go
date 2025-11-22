package entities

type User struct {
	ID   int
	TgId int64
}

type UserChannel struct {
	ID        int
	UserID    int
	ChannelID int
	Role      string
	Anonimity bool
}
