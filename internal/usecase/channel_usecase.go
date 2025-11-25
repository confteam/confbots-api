package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain"
)

type ChannelUseCase struct {
	rC domain.ChannelRepository
	rU domain.UserRepository
}

func NewChannelUseCase(rC domain.ChannelRepository, rU domain.UserRepository) *ChannelUseCase {
	return &ChannelUseCase{
		rC: rC,
		rU: rU,
	}
}

const channelPkg = "usecase.ChannelUseCase"

func (uc *ChannelUseCase) Create(
	ctx context.Context,
	channel domain.ChannelWithoutID,
) (int, error) {
	const op = channelPkg + ".Create"

	id, err := uc.rC.Create(ctx, channel)
	if err != nil {
		if errors.Is(err, domain.ErrChannelExists) {
			return 0, err
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, nil
}

func (uc *ChannelUseCase) Update(
	ctx context.Context,
	channel domain.ChannelWithoutCode,
) (*domain.Channel, error) {
	const op = channelPkg + ".Update"

	updatedChannel, err := uc.rC.Update(ctx, channel)
	if err != nil {
		if errors.Is(err, domain.ErrChannelNotFound) {
			return nil, err
		} else if errors.Is(err, domain.ErrChannelExists) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return updatedChannel, err
}

func (uc *ChannelUseCase) FindByCode(ctx context.Context, code string) (*domain.ChannelWithoutCode, error) {
	const op = channelPkg + ".FindByCode"

	channel, err := uc.rC.FindByCode(ctx, code)
	if err != nil {
		if errors.Is(err, domain.ErrChannelNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return channel, err
}

func (uc *ChannelUseCase) FindByID(ctx context.Context, id int) (*domain.ChannelWithoutID, error) {
	const op = channelPkg + ".FindByID"

	channel, err := uc.rC.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrChannelNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return channel, err
}

func (uc *ChannelUseCase) FindByChatID(ctx context.Context, chatID int64) (int, error) {
	const op = channelPkg + ".FindByChatID"

	id, err := uc.rC.FindByChatID(ctx, chatID)
	if err != nil {
		if errors.Is(err, domain.ErrChannelNotFound) {
			return 0, err
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, err
}

func (uc *ChannelUseCase) GetAllUserChannels(ctx context.Context, tgID int64) ([]domain.ChannelIDWithChannelChat, error) {
	const op = channelPkg + ".GetAllUserChannels"

	userID, err := uc.rU.GetIdByTgId(ctx, tgID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	channels, err := uc.rC.GetAllUserChannels(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserChannelNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return channels, nil
}
