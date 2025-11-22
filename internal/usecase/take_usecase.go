package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
	"github.com/jackc/pgx/v5"
)

type TakeUseCase struct {
	rU repositories.UserRepository
	rT repositories.TakeRepository
}

func NewTakeUseCase(rU repositories.UserRepository, rT repositories.TakeRepository) *TakeUseCase {
	return &TakeUseCase{
		rU: rU,
		rT: rT,
	}
}

const takePkg = "usecase.TakeUseCase"

func (uc *TakeUseCase) Create(ctx context.Context, tgid int64, userMessageID int64, adminMessageID int64, channelID int) (*entities.Take, error) {
	const op = takePkg + ".Create"

	userID, err := uc.rU.GetIdByTgId(ctx, tgid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	userChannelID, err := uc.rU.GetUserChannelID(ctx, userID, channelID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("userChannel not found")
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	take, err := uc.rT.Create(ctx, userMessageID, adminMessageID, userChannelID, channelID)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}

func (uc *TakeUseCase) GetById(ctx context.Context, id int, channelID int) (*entities.Take, error) {
	const op = takePkg + ".GetById"

	take, err := uc.rT.GetByID(ctx, id, channelID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("take not found")
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}

func (uc *TakeUseCase) GetByMsgId(ctx context.Context, messageID int64, channelID int) (*entities.Take, error) {
	const op = takePkg + ".GetByMsgId"

	take, err := uc.rT.GetByMsgID(ctx, messageID, channelID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}

func (uc *TakeUseCase) UpdateStatus(ctx context.Context, id int, channelID int) error {
	const op = takePkg + ".UpdateStatus"

	if err := uc.rT.UpdateStatus(ctx, id, channelID); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("take not found")
		}
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}

func (uc *TakeUseCase) GetTakeAuthor(ctx context.Context, id int, channelID int) (int64, error) {
	const op = takePkg + ".GetTakeAuthor"

	take, err := uc.rT.GetByID(ctx, id, channelID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("take not found")
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	userChannel, err := uc.rU.GetUserChannelByID(ctx, take.UserChannelID)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	tgid, err := uc.rU.GetTgIdById(ctx, userChannel.UserID)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return tgid, nil
}
