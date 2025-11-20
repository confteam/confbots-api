package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type BotRepository interface {
	FindBotByTgIdAndType(ctx context.Context, tgid int32, botType entities.BotType) (*entities.BotWithChannel, error)
	Create(ctx context.Context, tgid int32, botType entities.BotType) (*entities.BotWithChannel, error)
}
