package usecase

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type ChannelUseCase struct {
	r repositories.ChannelRepository
}

func NewChannelUseCase(r repositories.ChannelRepository) *ChannelUseCase {
	return &ChannelUseCase{
		r: r,
	}
}

const channelPkg = "usecase.ChannelUseCase"

func (uc *ChannelUseCase) Create(
	ctx context.Context,
	channel entities.ChannelWithoutID,
) (*entities.Channel, error) {
	const op = channelPkg + ".Create"

	newChannel, err := uc.r.Create(ctx, channel)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return newChannel, nil
}

func (uc *ChannelUseCase) Update(
	ctx context.Context,
	channel entities.ChannelWithoutIDAndCode,
) (*entities.Channel, error) {
	const op = channelPkg + ".Update"

	updatedChannel, err := uc.r.Update(ctx, channel)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return updatedChannel, err
}
