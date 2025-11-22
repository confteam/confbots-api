package repositories

import (
	"context"

	"github.com/confteam/confbots-api/internal/domain/entities"
)

type ReplyRepository interface {
	Create(ctx context.Context, userMessageID int64, adminMessageID int64, takeID int, channelID int) (int, error)
	GetByMsgId(ctx context.Context, messageID int64, channelID int) (*entities.Reply, error)
}
