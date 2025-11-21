package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
	"github.com/jackc/pgx/v5/pgtype"
)

type BotPostgresRepository struct {
	q *db.Queries
}

func NewBotPostgresRepository(q *db.Queries) repositories.BotRepository {
	return &BotPostgresRepository{
		q: q,
	}
}

const botPkg = "infrasctructure.repository.BotPostgresRepository"

func (r *BotPostgresRepository) FindBotByTgIdAndType(
	ctx context.Context,
	tgid int64,
	botType entities.BotType,
) (*entities.BotWithChannel, error) {
	const op = botPkg + "FindBotByTgIdAndType"

	botWithChannel, err := r.q.FindBotByTgIdAndType(ctx, db.FindBotByTgIdAndTypeParams{
		Tgid: tgid,
		Type: string(botType),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	var channel *entities.Channel
	if botWithChannel.ChannelID.Valid {
		channel = &entities.Channel{
			ID:                int(*ptrInt32(botWithChannel.ChannelID)),
			Code:              botWithChannel.Code.String,
			ChannelChatID:     ptrInt64(botWithChannel.ChannelChatID),
			AdminChatID:       ptrInt64(botWithChannel.AdminChatID),
			DiscussionsChatID: ptrInt64(botWithChannel.DiscussionsChatID),
			Decorations:       ptrString(botWithChannel.Decorations),
		}
	}

	return &entities.BotWithChannel{
		ID:      botWithChannel.ID,
		TgID:    botWithChannel.Tgid,
		Type:    entities.BotType(botWithChannel.Type),
		Channel: channel,
	}, nil
}

func (r *BotPostgresRepository) Create(
	ctx context.Context, tgid int64, botType entities.BotType,
) (*entities.BotWithChannel, error) {
	const op = botPkg + ".Create"

	bot, err := r.q.CreateBot(ctx, db.CreateBotParams{
		Tgid: tgid,
		Type: string(botType),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.BotWithChannel{
		ID:      bot.ID,
		TgID:    bot.Tgid,
		Type:    entities.BotType(bot.Type),
		Channel: nil,
	}, nil
}

func (r *BotPostgresRepository) UpdateChannelID(
	ctx context.Context, tgid int64, botType entities.BotType, channelID int,
) error {
	const op = botPkg + ".UpdateChannelID"

	var channelIDPgInt4 pgtype.Int4
	channelIDPgInt4.Int32 = int32(channelID)
	channelIDPgInt4.Valid = true

	if err := r.q.UpdateBotChannelID(ctx, db.UpdateBotChannelIDParams{
		Tgid:      tgid,
		Type:      string(botType),
		ChannelID: channelIDPgInt4,
	}); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}
