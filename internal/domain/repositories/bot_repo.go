package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type BotRepository interface {
	CreateIfNotExists(ctx context.Context, tgid int32, botType entities.BotType) (*entities.Bot, error)
}
