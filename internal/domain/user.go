package domain

import (
	"context"
)

type User struct {
	ID   int
	TgId int64
}

type UserChannel struct {
	ID        int
	UserID    int
	ChannelID int
	Role      string
	Anonimity bool
}

type UserRepository interface {
	Upsert(ctx context.Context, tgid int64, channelID int) (int, error)
	UpdateRole(ctx context.Context, role string, userID int, channelID int) error
	GetRole(ctx context.Context, userID int, channelID int) (string, error)
	GetIdByTgId(ctx context.Context, tgid int64) (int, error)
	GetTgIdById(ctx context.Context, id int) (int64, error)
	GetAnonimity(ctx context.Context, userID int, channelID int) (bool, error)
	ToggleAnonimity(ctx context.Context, userID int, channelID int) (bool, error)
	GetUserChannelID(ctx context.Context, userID int, channelID int) (int, error)
	GetUserChannelByID(ctx context.Context, id int) (*UserChannel, error)
	GetAllUsersInChannel(ctx context.Context, channelID int) ([]int64, error)
}
