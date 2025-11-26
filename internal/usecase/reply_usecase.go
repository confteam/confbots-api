package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain"
)

type ReplyUseCase struct {
	r domain.ReplyRepository
}

func NewReplyUseCase(r domain.ReplyRepository) *ReplyUseCase {
	return &ReplyUseCase{
		r: r,
	}
}

const replyPkg = "usecase.ReplyUseCase"

func (uc *ReplyUseCase) Create(
	ctx context.Context,
	userMessageID int64,
	adminMessageID int64,
	takeID int,
	channelID int,
) (int, error) {
	const op = replyPkg + ".Create"

	id, err := uc.r.Create(ctx, userMessageID, adminMessageID, takeID, channelID)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, nil
}

func (uc *ReplyUseCase) GetByMsgIDAndChannelID(
	ctx context.Context,
	messageID int64,
	channelID int,
) (*domain.Reply, error) {
	const op = replyPkg + ".GetByMsgIDAndChannelID"

	reply, err := uc.r.GetByMsgIDAndChannelID(ctx, messageID, channelID)
	if err != nil {
		if errors.Is(err, domain.ErrReplyNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return reply, nil
}

func (uc *ReplyUseCase) GetByMsgIDAndTakeID(
	ctx context.Context,
	messageID int64,
	takeID int,
) (*domain.Reply, error) {
	const op = replyPkg + ".GetByMsgIDAndTakeID"

	reply, err := uc.r.GetByMsgIDAndTakeID(ctx, messageID, takeID)
	if err != nil {
		if errors.Is(err, domain.ErrReplyNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return reply, nil
}
