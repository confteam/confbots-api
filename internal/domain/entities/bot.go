package entities

type BotType string

const (
	BotTypeTakes BotType = "TAKES"
	BotTypeMP    BotType = "MP"
	BotTypeMod   BotType = "MOD"
)

type BotWithChannel struct {
	ID   int32
	TgID int64
	Type BotType

	Channel *Channel
}
