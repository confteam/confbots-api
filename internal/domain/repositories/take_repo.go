package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type TakeRepository interface {
	Create(ctx context.Context, userMessageID int64, adminMessageID int64, userChannelID int, channelID int) (*entities.Take, error)
	GetByID(ctx context.Context, id int, channelID int) (*entities.Take, error)
	GetByMsgID(ctx context.Context, messageID int64, channelID int) (*entities.Take, error)
}
