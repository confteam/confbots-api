package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain"
)

type UserUseCase struct {
	r domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		r: r,
	}
}

const userPkg = "usecase.UserUseCase"

func (uc *UserUseCase) Upsert(
	ctx context.Context,
	tgid int64,
	channelID int,
) (int, error) {
	const op = userPkg + ".Upsert"

	id, err := uc.r.Upsert(ctx, tgid, channelID)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, nil
}

func (uc *UserUseCase) UpdateRole(
	ctx context.Context,
	role string,
	tgid int64,
	channelID int,
) error {
	const op = userPkg + ".UpdateRole"

	userID, err := uc.r.GetIdByTgId(ctx, tgid)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return err
		}
		return fmt.Errorf("%s:%v", op, err)
	}

	if err := uc.r.UpdateRole(ctx, role, userID, channelID); err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return err
		}
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (uc *UserUseCase) GetRole(
	ctx context.Context,
	tgid int64,
	channelID int,
) (string, error) {
	const op = userPkg + ".GetRole"

	userID, err := uc.r.GetIdByTgId(ctx, tgid)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return "", err
		}
		return "", fmt.Errorf("%s:%v", op, err)
	}

	role, err := uc.r.GetRole(ctx, userID, channelID)
	if err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return "", err
		}
		return "", fmt.Errorf("%s:%v", op, err)
	}

	return role, nil
}

func (uc *UserUseCase) GetAnonimity(
	ctx context.Context,
	tgid int64,
	channelID int,
) (bool, error) {
	const op = userPkg + ".GetAnonimity"

	userID, err := uc.r.GetIdByTgId(ctx, tgid)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return false, err
		}
		return false, fmt.Errorf("%s:%v", op, err)
	}

	anonimity, err := uc.r.GetAnonimity(ctx, userID, channelID)
	if err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return false, err
		}
		return false, fmt.Errorf("%s:%v", op, err)
	}

	return anonimity, nil
}

func (uc *UserUseCase) ToggleAnonimity(
	ctx context.Context,
	tgid int64,
	channelID int,
) (bool, error) {
	const op = userPkg + ".GetAnonimity"

	userID, err := uc.r.GetIdByTgId(ctx, tgid)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return false, err
		}
		return false, fmt.Errorf("%s:%v", op, err)
	}

	anonimity, err := uc.r.ToggleAnonimity(ctx, userID, channelID)
	if err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return false, err
		}
		return false, fmt.Errorf("%s:%v", op, err)
	}

	return anonimity, nil
}
