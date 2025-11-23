package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

type ReplyPostgresRepository struct {
	q *db.Queries
}

func NewReplyPostgresRepository(q *db.Queries) domain.ReplyRepository {
	return &ReplyPostgresRepository{
		q: q,
	}
}

const replyPkg = "infrastructure.repository.ReplyPostgresRepository"

func (r *ReplyPostgresRepository) Create(
	ctx context.Context,
	userMessageID int64,
	adminMessageID int64,
	takeID int,
	channelID int,
) (int, error) {
	const op = replyPkg + ".Create"

	id, err := r.q.CreateReply(ctx, db.CreateReplyParams{
		UserMessageID:  userMessageID,
		AdminMessageID: adminMessageID,
		TakeID:         int32(takeID),
		ChannelID:      int32(channelID),
	})
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *ReplyPostgresRepository) GetByMsgId(
	ctx context.Context,
	messageID int64,
	channelID int,
) (*domain.Reply, error) {
	const op = replyPkg + ".GetByMsgId"

	reply, err := r.q.GetReplyByMsgId(ctx, db.GetReplyByMsgIdParams{
		UserMessageID: messageID,
		ChannelID:     int32(channelID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrReplyNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.Reply{
		ID:             int(reply.ID),
		UserMessageID:  reply.UserMessageID,
		AdminMessageID: reply.AdminMessageID,
		TakeID:         int(reply.TakeID),
		ChannelID:      int(reply.ChannelID),
	}, nil
}
