package repository

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type ReplyPostgresRepository struct {
	q *db.Queries
}

func NewReplyPostgresRepository(q *db.Queries) repositories.ReplyRepository {
	return &ReplyPostgresRepository{
		q: q,
	}
}

const replyPkg = "infrastructure.repository.ReplyPostgresRepository"

func (r *ReplyPostgresRepository) Create(ctx context.Context, userMessageID int64, adminMessageID int64, takeID int) (int, error) {
	const op = replyPkg + ".Create"

	id, err := r.q.CreateReply(ctx, db.CreateReplyParams{
		UserMessageID:  userMessageID,
		AdminMessageID: adminMessageID,
		TakeID:         int32(takeID),
	})
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *ReplyPostgresRepository) GetByMsgId(ctx context.Context, messageID int64, takeID int) (*entities.Reply, error) {
	const op = replyPkg + ".GetByMsgId"

	reply, err := r.q.GetReplyByMsgId(ctx, db.GetReplyByMsgIdParams{
		UserMessageID: messageID,
		TakeID:        int32(takeID),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.Reply{
		ID:             int(reply.ID),
		UserMessageID:  reply.UserMessageID,
		AdminMessageID: reply.AdminMessageID,
		TakeID:         int(reply.TakeID),
	}, nil
}
