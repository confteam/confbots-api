package repository

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
	"github.com/jackc/pgx/v5"
)

type TakePostgresRepository struct {
	q *db.Queries
}

func NewTakePostgresRepository(q *db.Queries) repositories.TakeRepository {
	return &TakePostgresRepository{
		q: q,
	}
}

const takePkg = "infrastructure.repository.TakePostgresRepository"

func (r *TakePostgresRepository) Create(ctx context.Context, userMessageID int64, adminMessageID int64, userChannelID int, channelID int) (*entities.Take, error) {
	const op = takePkg + ".Create"

	take, err := r.q.CreateTake(ctx, db.CreateTakeParams{
		UserMessageID:  userMessageID,
		AdminMessageID: adminMessageID,
		UserChannelID:  int32(userChannelID),
		ChannelID:      int32(channelID),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.Take{
		ID:             int(take.ID),
		UserMessageID:  take.UserMessageID,
		AdminMessageID: take.AdminMessageID,
		UserChannelID:  int(take.UserChannelID),
		ChannelID:      int(take.ChannelID),
	}, nil
}

func (r *TakePostgresRepository) GetByID(ctx context.Context, id int, channelID int) (*entities.Take, error) {
	const op = takePkg + ".GetByID"

	take, err := r.q.GetTakeById(ctx, db.GetTakeByIdParams{
		ID:        int32(id),
		ChannelID: int32(channelID),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.Take{
		ID:             int(take.ID),
		Status:         take.Status,
		UserMessageID:  take.UserMessageID,
		AdminMessageID: take.AdminMessageID,
		UserChannelID:  int(take.UserChannelID),
		ChannelID:      int(take.ChannelID),
	}, nil
}

func (r *TakePostgresRepository) GetByMsgID(ctx context.Context, messageID int64, channelID int) (*entities.Take, error) {
	const op = takePkg + ".GetByMsgID"

	take, err := r.q.GetTakeByMsgId(ctx, db.GetTakeByMsgIdParams{
		UserMessageID: messageID,
		ChannelID:     int32(channelID),
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.Take{
		ID:             int(take.ID),
		Status:         take.Status,
		UserMessageID:  take.UserMessageID,
		AdminMessageID: take.AdminMessageID,
		UserChannelID:  int(take.UserChannelID),
		ChannelID:      int(take.ChannelID),
	}, nil
}

func (r *TakePostgresRepository) UpdateStatus(ctx context.Context, id int, channelID int) error {
	const op = takePkg + ".UpdateStatus"

	if err := r.q.UpdateTakeStatus(ctx, db.UpdateTakeStatusParams{
		ID:        int32(id),
		ChannelID: int32(channelID),
	}); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}
