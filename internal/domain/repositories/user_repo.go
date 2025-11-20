package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type UserRepository interface {
	Upsert(ctx context.Context, tgid int64, channelID int) (int, error)
	UpdateRole(ctx context.Context, role entities.Role, userID int, channelID int) error
	GetRole(ctx context.Context, userID int, channelID int) (entities.Role, error)
}
