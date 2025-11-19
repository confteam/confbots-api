package entities

type BotType string

const (
	BotTypeTakes BotType = "TAKES"
	BotTypeMP    BotType = "MP"
	BotTypeMod   BotType = "MOD"
)

type Bot struct {
	ID        int32   `json:"id"`
	TgID      int32   `json:"tgid"`
	Type      BotType `json:"type"`
	ChannelID *int32  `json:"channel_id"`
}
