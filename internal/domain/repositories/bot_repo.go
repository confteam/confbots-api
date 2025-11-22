package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type BotRepository interface {
	FindBotByTgIdAndType(ctx context.Context, tgid int64, botType string) (*entities.BotWithChannel, error)
	Create(ctx context.Context, tgid int64, botType string) (*entities.BotWithChannel, error)
}
