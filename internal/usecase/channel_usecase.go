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
	channel domain.ChannelWithBotTgidAndType,
) (*domain.Channel, error) {
	const op = channelPkg + ".Create"

	newChannel, err := uc.r.Create(ctx, channel)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return newChannel, nil
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
