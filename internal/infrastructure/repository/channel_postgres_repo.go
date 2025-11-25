package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
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
	channel domain.ChannelWithoutID,
) (int, error) {
	const op = channelPkg + ".Create"

	id, err := r.q.CreateChannel(ctx, db.CreateChannelParams{
		Code:              channel.Code,
		ChannelChatID:     ptrPgInt8(channel.ChannelChatID),
		AdminChatID:       ptrPgInt8(channel.AdminChatID),
		DiscussionsChatID: ptrPgInt8(channel.DiscussionsChatID),
	})
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return 0, domain.ErrChannelExists
			}
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
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
		} else if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return nil, domain.ErrChannelExists
			}
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

func (r *ChannelPostgresRepository) FindByCode(ctx context.Context, code string) (*domain.ChannelWithoutCode, error) {
	const op = channelPkg + ".FindByCode"

	channel, err := r.q.FindChannelByCode(ctx, code)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrChannelNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.ChannelWithoutCode{
		ID:                int(channel.ID),
		ChannelChatID:     &channel.ChannelChatID.Int64,
		AdminChatID:       &channel.AdminChatID.Int64,
		DiscussionsChatID: &channel.DiscussionsChatID.Int64,
		Decorations:       &channel.Decorations.String,
	}, nil
}

func (r *ChannelPostgresRepository) FindByID(ctx context.Context, id int) (*domain.ChannelWithoutID, error) {
	const op = channelPkg + ".FindById"

	channel, err := r.q.FindChannelById(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrChannelNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &domain.ChannelWithoutID{
		Code:              channel.Code,
		ChannelChatID:     &channel.ChannelChatID.Int64,
		AdminChatID:       &channel.AdminChatID.Int64,
		DiscussionsChatID: &channel.DiscussionsChatID.Int64,
		Decorations:       &channel.Decorations.String,
	}, nil
}

func (r *ChannelPostgresRepository) FindByChatID(ctx context.Context, chatID int64) (int, error) {
	const op = channelPkg + ".FindByChatID"

	id, err := r.q.FindChannelByChatId(ctx, pgtype.Int8{Int64: chatID, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, domain.ErrChannelNotFound
		}
		return 0, fmt.Errorf("%s:%v", op, err)
	}

	return int(id), nil
}

func (r *ChannelPostgresRepository) GetAllUserChannels(ctx context.Context, userID int) ([]domain.ChannelIDWithChannelChat, error) {
	const op = channelPkg + ".GetAllUserChannels"

	channelsRaw, err := r.q.GetAllUserChannels(ctx, int32(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserChannelNotFound
		}
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	channels := make([]domain.ChannelIDWithChannelChat, len(channelsRaw))
	for i, v := range channelsRaw {
		channels[i] = domain.ChannelIDWithChannelChat{
			ID:            int(v.ChannelID),
			ChannelChatID: v.ChannelChatID.Int64,
		}
	}

	return channels, nil
}
