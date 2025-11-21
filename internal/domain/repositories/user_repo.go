package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type UserRepository interface {
	Upsert(ctx context.Context, tgid int64, channelID int, role entities.Role) (int, error)
	UpdateRole(ctx context.Context, role entities.Role, userID int, channelID int) error
	GetRole(ctx context.Context, userID int, channelID int) (entities.Role, error)
	GetIdByTgId(ctx context.Context, tgid int64) (int, error)
	GetAnonimity(ctx context.Context, userID int, channelID int) (bool, error)
	ToggleAnonimity(ctx context.Context, userID int, channelID int) (bool, error)
}
