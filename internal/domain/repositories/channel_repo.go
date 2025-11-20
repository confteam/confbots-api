package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type ChannelRepository interface {
	Create(ctx context.Context, channel entities.ChannelWithoutIDAndCode) (*entities.Channel, error)
	Update(ctx context.Context, channel entities.ChannelWithoutIDAndCode) (*entities.Channel, error)
}
