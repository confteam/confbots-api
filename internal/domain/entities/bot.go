package entities

type BotType string

const (
	BotTypeTakes BotType = "TAKES"
	BotTypeMP    BotType = "MP"
	BotTypeMod   BotType = "MOD"
)

type Bot struct {
	ID        int32
	TgID      int32
	Type      BotType
	ChannelID *int32
}
