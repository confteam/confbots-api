package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
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
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("take not found")
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return take, nil
}
