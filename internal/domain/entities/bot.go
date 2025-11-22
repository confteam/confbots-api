package entities

type BotWithChannel struct {
	ID   int32
	TgID int64
	Type string

	Channel *Channel
}
