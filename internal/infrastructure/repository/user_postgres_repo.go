package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserPostgresRepository struct {
	q *db.Queries
}

func NewUserPostgresRepository(q *db.Queries) domain.UserRepository {
	return &UserPostgresRepository{
		q: q,
	}
}

const userPkg = "infrasctructure.repository.UserPostgresRepository"

func (r *UserPostgresRepository) Upsert(
	ctx context.Context,
	tgid int64,
	channelID int,
) (int, error) {
	const op = userPkg + ".Upsert"

	id, err := r.q.UpsertUser(ctx, db.UpsertUserParams{
		Tgid:      tgid,
		ChannelID: int32(channelID),
	})
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *UserPostgresRepository) UpdateRole(
	ctx context.Context,
	role string,
	userID int,
	channelID int,
) error {
	const op = userPkg + ".UpdateRole"

	if err := r.q.UpdateUserRole(ctx, db.UpdateUserRoleParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
		Role:      string(role),
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrUserChannelNotFound
		}
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (r *UserPostgresRepository) GetRole(ctx context.Context, userID int, channelID int) (string, error) {
	const op = userPkg + ".GetRole"

	role, err := r.q.GetUserRole(ctx, db.GetUserRoleParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", domain.ErrUserChannelNotFound
		}
		return "", fmt.Errorf("%s:%v", op, err)
	}

	return role, nil
}

func (r *UserPostgresRepository) GetIdByTgId(ctx context.Context, tgid int64) (int, error) {
	const op = userPkg + ".GetIdByTgId"

	id, err := r.q.GetUserIdByTgId(ctx, tgid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, domain.ErrUserNotFound
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *UserPostgresRepository) GetTgIdById(ctx context.Context, id int) (int64, error) {
	const op = userPkg + ".GetTgIdById"

	tgid, err := r.q.GetUserTgIdById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, domain.ErrUserNotFound
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return tgid, nil
}

func (r *UserPostgresRepository) GetAnonimity(ctx context.Context, userID int, channelID int) (bool, error) {
	const op = userPkg + ".GetAnonimity"

	anonimity, err := r.q.GetUserAnonimity(ctx, db.GetUserAnonimityParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, domain.ErrUserNotFound
		}
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
		if errors.Is(err, pgx.ErrNoRows) {
			return false, domain.ErrUserNotFound
		}
		return false, fmt.Errorf("%s:%v", op, err)
	}

	return anonimity.Bool, nil
}

func (r *UserPostgresRepository) GetUserChannelID(ctx context.Context, userID int, channelID int) (int, error) {
	const op = userPkg + ".GetUserChannel"

	id, err := r.q.GetUserChannelId(ctx, db.GetUserChannelIdParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, domain.ErrUserChannelNotFound
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *UserPostgresRepository) GetUserChannelByID(ctx context.Context, id int) (*domain.UserChannel, error) {
	const op = userPkg + ".GetUserChannel"

	uc, err := r.q.GetUserChannelById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserChannelNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.UserChannel{
		ID:        int(uc.ID),
		UserID:    int(uc.UserID),
		ChannelID: int(uc.ChannelID),
		Role:      uc.Role,
		Anonimity: uc.Anonimity.Bool,
	}, nil
}
