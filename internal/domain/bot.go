package domain

import "context"

type BotWithChannel struct {
	ID   int32
	TgID int64
	Type string

	Channel *Channel
}

type BotRepository interface {
	FindBotByTgIdAndType(ctx context.Context, tgid int64, botType string) (*BotWithChannel, error)
	Create(ctx context.Context, tgid int64, botType string) (*BotWithChannel, error)
}
