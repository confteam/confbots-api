package usecase

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/internal/domain"
)

type BotUseCase struct {
	r domain.BotRepository
}

func NewBotUseCase(repo domain.BotRepository) *BotUseCase {
	return &BotUseCase{
		r: repo,
	}
}

const botPkg = "usecase.BotUseCase"

func (uc *BotUseCase) Auth(
	ctx context.Context,
	tgid int64,
	botType string,
) (*domain.BotWithChannel, error) {
	const op = botPkg + ".Auth"
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
