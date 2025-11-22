package usecase

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type ReplyUseCase struct {
	r repositories.ReplyRepository
}

func NewReplyUseCase(r repositories.ReplyRepository) *ReplyUseCase {
	return &ReplyUseCase{
		r: r,
	}
}

const replyPkg = "usecase.ReplyUseCase"

func (uc *ReplyUseCase) Create(ctx context.Context, userMessageID int64, adminMessageID int64, takeID int, channelID int) (int, error) {
	const op = replyPkg + ".Create"

	id, err := uc.r.Create(ctx, userMessageID, adminMessageID, takeID, channelID)
	if err != nil {
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return id, nil
}

func (uc *ReplyUseCase) GetByMsgID(ctx context.Context, messageID int64, takeID int, channelID int) (*entities.Reply, error) {
	const op = replyPkg + ".GetByMsgID"

	reply, err := uc.r.GetByMsgId(ctx, messageID, takeID, channelID)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return reply, nil
}
