package usecase

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type BotUseCase struct {
	r repositories.BotRepository
}

func NewBotUseCase(repo repositories.BotRepository) *BotUseCase {
	return &BotUseCase{
		r: repo,
	}
}

const pkg = "usecase.BotUseCase"

func (uc *BotUseCase) CreateIfNotExists(
	ctx context.Context,
	tgid int32,
	botType entities.BotType,
) (*entities.Bot, error) {
	const op = pkg + ".CreateIfNotExists"
	bot, err := uc.r.CreateIfNotExists(ctx, tgid, botType)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return bot, nil
}
