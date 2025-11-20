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

func (uc *BotUseCase) Auth(
	ctx context.Context,
	tgid int32,
	botType entities.BotType,
) (*entities.BotWithChannel, error) {
	const op = pkg + ".Auth"
	bot, err := uc.r.FindBotByTgIdAndType(ctx, tgid, botType)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	if bot != nil {
		return bot, nil
	}

	newBot, err := uc.r.Create(ctx, tgid, botType)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return newBot, nil
}
