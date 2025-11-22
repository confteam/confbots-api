package repositories

import (
	"context"
)

type UserRepository interface {
	Upsert(ctx context.Context, tgid int64, channelID int, role string) (int, error)
	UpdateRole(ctx context.Context, role string, userID int, channelID int) error
	GetRole(ctx context.Context, userID int, channelID int) (string, error)
	GetIdByTgId(ctx context.Context, tgid int64) (int, error)
	GetAnonimity(ctx context.Context, userID int, channelID int) (bool, error)
	ToggleAnonimity(ctx context.Context, userID int, channelID int) (bool, error)
	GetUserChannelID(ctx context.Context, userID int, channelID int) (int, error)
}
