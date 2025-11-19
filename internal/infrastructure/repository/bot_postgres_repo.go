package repository

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type BotPostgresRepository struct {
	q *db.Queries
}

func NewBotPostgresRepository(q *db.Queries) repositories.BotRepository {
	return &BotPostgresRepository{
		q: q,
	}
}

const pkg = "infrasctructure.repository"

func (r *BotPostgresRepository) CreateIfNotExists(
	ctx context.Context, tgid int32, botType entities.BotType,
) (*entities.Bot, error) {
	const op = pkg + ".CreateIfNotExists"
	bot, err := r.q.CreateIfNotExists(ctx, db.CreateIfNotExistsParams{
		Tgid: tgid,
		Type: string(botType),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	var channelID *int32
	if bot.ChannelID.Valid {
		channelID = &bot.ChannelID.Int32
	}

	return &entities.Bot{
		ID:        bot.ID,
		TgID:      bot.Tgid,
		Type:      entities.BotType(bot.Type),
		ChannelID: channelID,
	}, nil
}
