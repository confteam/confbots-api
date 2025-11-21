package usecase

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type UserUseCase struct {
	r repositories.UserRepository
}

func NewUserUseCase(r repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		r: r,
	}
}

const userPkg = "usecase.UserUseCase"

func (uc *UserUseCase) Upsert(ctx context.Context, tgid int64, channelID int) (int, error) {
	const op = userPkg + ".Upsert"

	id, err := uc.r.Upsert(ctx, tgid, channelID)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, nil
}

func (uc *UserUseCase) UpdateRole(ctx context.Context, role entities.Role, tgid int64, channelID int) error {
	const op = userPkg + ".UpdateRole"

	userID, err := uc.r.GetIdByTgId(ctx, tgid)
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	if err := uc.r.UpdateRole(ctx, role, userID, channelID); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (uc *UserUseCase) GetRole(ctx context.Context, tgid int64, channelID int) (entities.Role, error) {
	const op = userPkg + ".GetRole"

	userID, err := uc.r.GetIdByTgId(ctx, tgid)
	if err != nil {
		return "", fmt.Errorf("%s:%v", op, err)
	}

	role, err := uc.r.GetRole(ctx, userID, channelID)
	if err != nil {
		return "", fmt.Errorf("%s:%v", op, err)
	}

	return entities.Role(role), nil
}
