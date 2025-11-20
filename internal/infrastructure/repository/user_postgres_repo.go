package repository

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
	"github.com/jackc/pgx/v5/pgtype"
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

func (r *UserPostgresRepository) Upsert(ctx context.Context, tgid int64, channelID int) (int, error) {
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

func (r *UserPostgresRepository) UpdateRole(ctx context.Context, role entities.Role, userID int, channelID int) error {
	const op = userPkg + ".UpdateRole"

	var pgRole pgtype.Text
	pgRole.String = string(role)
	pgRole.Valid = true

	if err := r.q.UpdateUserRole(ctx, db.UpdateUserRoleParams{
		UserID:    int32(userID),
		ChannelID: int32(channelID),
		Role:      pgRole,
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

	return entities.Role(*ptrString(role)), nil
}
