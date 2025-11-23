package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

type BotPostgresRepository struct {
	q *db.Queries
}

func NewBotPostgresRepository(q *db.Queries) domain.BotRepository {
	return &BotPostgresRepository{
		q: q,
	}
}

const botPkg = "infrasctructure.repository.BotPostgresRepository"

func (r *BotPostgresRepository) FindBotByTgIdAndType(
	ctx context.Context,
	tgid int64,
	botType string,
) (*domain.BotWithChannel, error) {
	const op = botPkg + "FindBotByTgIdAndType"

	botWithChannel, err := r.q.FindBotByTgIdAndType(ctx, db.FindBotByTgIdAndTypeParams{
		Tgid: tgid,
		Type: string(botType),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrBotNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	var channel *domain.Channel
	if botWithChannel.ChannelID.Valid {
		channel = &domain.Channel{
			ID:                int(*ptrInt32(botWithChannel.ChannelID)),
			Code:              botWithChannel.Code.String,
			ChannelChatID:     ptrInt64(botWithChannel.ChannelChatID),
			AdminChatID:       ptrInt64(botWithChannel.AdminChatID),
			DiscussionsChatID: ptrInt64(botWithChannel.DiscussionsChatID),
			Decorations:       ptrString(botWithChannel.Decorations),
		}
	}

	return &domain.BotWithChannel{
		ID:      botWithChannel.ID,
		TgID:    botWithChannel.Tgid,
		Type:    botWithChannel.Type,
		Channel: channel,
	}, nil
}

func (r *BotPostgresRepository) Create(
	ctx context.Context, tgid int64, botType string,
) (*domain.BotWithChannel, error) {
	const op = botPkg + ".Create"

	bot, err := r.q.CreateBot(ctx, db.CreateBotParams{
		Tgid: tgid,
		Type: string(botType),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.BotWithChannel{
		ID:      bot.ID,
		TgID:    bot.Tgid,
		Type:    bot.Type,
		Channel: nil,
	}, nil
}
