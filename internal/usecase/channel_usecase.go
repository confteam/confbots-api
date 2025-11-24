package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain"
)

type ChannelUseCase struct {
	r domain.ChannelRepository
}

func NewChannelUseCase(r domain.ChannelRepository) *ChannelUseCase {
	return &ChannelUseCase{
		r: r,
	}
}

const channelPkg = "usecase.ChannelUseCase"

func (uc *ChannelUseCase) Create(
	ctx context.Context,
	channel domain.ChannelWithoutID,
) (int, error) {
	const op = channelPkg + ".Create"

	id, err := uc.r.Create(ctx, channel)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, nil
}

func (uc *ChannelUseCase) Update(
	ctx context.Context,
	channel domain.ChannelWithoutCode,
) (*domain.Channel, error) {
	const op = channelPkg + ".Update"

	updatedChannel, err := uc.r.Update(ctx, channel)
	if err != nil {
		if errors.Is(err, domain.ErrChannelNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return updatedChannel, err
}

func (uc *ChannelUseCase) FindByCode(ctx context.Context, code string) (*domain.ChannelWithoutCode, error) {
	const op = channelPkg + ".FindByCode"

	channel, err := uc.r.FindByCode(ctx, code)
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

	channel, err := uc.r.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrChannelNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return channel, err
}
