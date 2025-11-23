package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain"
)

type TakeUseCase struct {
	rU domain.UserRepository
	rT domain.TakeRepository
}

func NewTakeUseCase(rU domain.UserRepository, rT domain.TakeRepository) *TakeUseCase {
	return &TakeUseCase{
		rU: rU,
		rT: rT,
	}
}

const takePkg = "usecase.TakeUseCase"

func (uc *TakeUseCase) Create(
	ctx context.Context,
	tgid int64,
	userMessageID int64,
	adminMessageID int64,
	channelID int,
) (*domain.Take, error) {
	const op = takePkg + ".Create"

	userID, err := uc.rU.GetIdByTgId(ctx, tgid)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	userChannelID, err := uc.rU.GetUserChannelID(ctx, userID, channelID)
	if err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	take, err := uc.rT.Create(ctx, userMessageID, adminMessageID, userChannelID, channelID)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}

func (uc *TakeUseCase) GetById(ctx context.Context, id int) (*domain.Take, error) {
	const op = takePkg + ".GetById"

	take, err := uc.rT.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrTakeNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}

func (uc *TakeUseCase) GetByMsgId(
	ctx context.Context,
	messageID int64,
	channelID int,
) (*domain.Take, error) {
	const op = takePkg + ".GetByMsgId"

	take, err := uc.rT.GetByMsgID(ctx, messageID, channelID)
	if err != nil {
		if errors.Is(err, domain.ErrTakeNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}

func (uc *TakeUseCase) UpdateStatus(ctx context.Context, id int, status string) error {
	const op = takePkg + ".UpdateStatus"

	if err := uc.rT.UpdateStatus(ctx, id, status); err != nil {
		if errors.Is(err, domain.ErrTakeNotFound) {
			return err
		}
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (uc *TakeUseCase) GetTakeAuthor(ctx context.Context, id int) (int64, error) {
	const op = takePkg + ".GetTakeAuthor"

	take, err := uc.rT.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrTakeNotFound) {
			return 0, err
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	userChannel, err := uc.rU.GetUserChannelByID(ctx, take.UserChannelID)
	if err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return 0, err
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	tgid, err := uc.rU.GetTgIdById(ctx, userChannel.UserID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return 0, err
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return tgid, nil
}
