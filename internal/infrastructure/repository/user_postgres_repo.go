package repository

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type UserPostgresRepository struct {
	q *db.Queries
}

func NewUserPostgresRepository(q *db.Queries) repositories.UserRepository {
	return &UserPostgresRepository{
		q: q,
	}
}

const userPkg = "infrasctructure.repository.UserPostgresRepository"

func (r *UserPostgresRepository) Upsert(ctx context.Context, tgid int64, channelID int, role entities.Role) (int, error) {
	const op = userPkg + ".Upsert"

	id, err := r.q.UpsertUser(ctx, db.UpsertUserParams{
		Role:      string(role),
		Tgid:      tgid,
		ChannelID: int32(channelID),
	})
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *UserPostgresRepository) UpdateRole(ctx context.Context, role entities.Role, userID int, channelID int) error {
	const op = userPkg + ".UpdateRole"

	if err := r.q.UpdateUserRole(ctx, db.UpdateUserRoleParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
		Role:      string(role),
	}); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (r *UserPostgresRepository) GetRole(ctx context.Context, userID int, channelID int) (entities.Role, error) {
	const op = userPkg + ".GetRole"

	role, err := r.q.GetUserRole(ctx, db.GetUserRoleParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
	})
	if err != nil {
		return "", fmt.Errorf("%s:%v", op, err)
	}

	return entities.Role(role), nil
}

func (r *UserPostgresRepository) GetIdByTgId(ctx context.Context, tgid int64) (int, error) {
	const op = userPkg + ".GetIdByTgId"

	id, err := r.q.GetUserIdByTgId(ctx, tgid)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *UserPostgresRepository) GetAnonimity(ctx context.Context, userID int, channelID int) (bool, error) {
	const op = userPkg + ".GetAnonimity"

	anonimity, err := r.q.GetUserAnonimity(ctx, db.GetUserAnonimityParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
	})
	if err != nil {
		return false, fmt.Errorf("%s:%v", op, err)
	}

	return anonimity.Bool, nil
}

func (r *UserPostgresRepository) ToggleAnonimity(ctx context.Context, userID int, channelID int) (bool, error) {
	const op = userPkg + ".ToggleAnonimity"

	anonimity, err := r.q.ToggleUserAnonimity(ctx, db.ToggleUserAnonimityParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
	})
	if err != nil {
		return false, fmt.Errorf("%s:%v", op, err)
	}

	return anonimity.Bool, nil
}
