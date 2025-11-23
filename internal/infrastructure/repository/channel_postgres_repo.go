package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain"
	"github.com/jackc/pgx/v5"
)

type ChannelPostgresRepository struct {
	q *db.Queries
}

func NewChannelPostgresRepository(q *db.Queries) domain.ChannelRepository {
	return &ChannelPostgresRepository{
		q: q,
	}
}

const channelPkg = "infrastructure.repository.ChannelPostgresRepository"

func (r *ChannelPostgresRepository) Create(
	ctx context.Context,
	channel domain.ChannelWithBotTgidAndType,
) (*domain.Channel, error) {
	const op = channelPkg + ".Create"

	newChannel, err := r.q.CreateChannel(ctx, db.CreateChannelParams{
		Tgid:              int64(channel.BotTgID),
		Type:              channel.BotType,
		Code:              channel.Code,
		ChannelChatID:     ptrPgInt8(channel.ChannelChatID),
		AdminChatID:       ptrPgInt8(channel.AdminChatID),
		DiscussionsChatID: ptrPgInt8(channel.DiscussionsChatID),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.Channel{
		ID:                int(newChannel.ID),
		Code:              newChannel.Code,
		ChannelChatID:     &newChannel.ChannelChatID.Int64,
		AdminChatID:       &newChannel.AdminChatID.Int64,
		DiscussionsChatID: &newChannel.DiscussionsChatID.Int64,
		Decorations:       &newChannel.Decorations.String,
	}, nil
}

func (r *ChannelPostgresRepository) Update(
	ctx context.Context,
	channel domain.ChannelWithoutCode,
) (*domain.Channel, error) {
	const op = channelPkg + ".Update"

	updatedChannel, err := r.q.UpdateChannel(ctx, db.UpdateChannelParams{
		ID:                int32(channel.ID),
		ChannelChatID:     ptrPgInt8(channel.ChannelChatID),
		AdminChatID:       ptrPgInt8(channel.AdminChatID),
		DiscussionsChatID: ptrPgInt8(channel.DiscussionsChatID),
		Decorations:       ptrPgText(channel.Decorations),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrChannelNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.Channel{
		ID:                channel.ID,
		Code:              updatedChannel.Code,
		ChannelChatID:     &updatedChannel.ChannelChatID.Int64,
		AdminChatID:       &updatedChannel.AdminChatID.Int64,
		DiscussionsChatID: &updatedChannel.DiscussionsChatID.Int64,
		Decorations:       &updatedChannel.Decorations.String,
	}, nil
}
