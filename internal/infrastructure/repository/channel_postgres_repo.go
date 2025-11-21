package repository

import (
	"context"
	"fmt"

	"github.com/confteam/confbots-api/db"
	"github.com/confteam/confbots-api/internal/domain/entities"
	"github.com/confteam/confbots-api/internal/domain/repositories"
)

type ChannelPostgresRepository struct {
	q *db.Queries
}

func NewChannelPostgresRepository(q *db.Queries) repositories.ChannelRepository {
	return &ChannelPostgresRepository{
		q: q,
	}
}

const channelPkg = "infrastructure.repository.ChannelPostgresRepository"

func (r *ChannelPostgresRepository) Create(
	ctx context.Context,
	channel entities.ChannelWithoutID,
) (*entities.Channel, error) {
	const op = channelPkg + ".Create"

	newChannel, err := r.q.CreateChannel(ctx, db.CreateChannelParams{
		Code:              channel.Code,
		ChannelChatID:     ptrPgInt8(channel.ChannelChatID),
		AdminChatID:       ptrPgInt8(channel.AdminChatID),
		DiscussionsChatID: ptrPgInt8(channel.DiscussionsChatID),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.Channel{
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
	channel entities.ChannelWithoutIDAndCode,
) (*entities.Channel, error) {
	const op = channelPkg + ".Update"

	updatedChannel, err := r.q.UpdateChannel(ctx, db.UpdateChannelParams{
		ChannelChatID:     ptrPgInt8(channel.ChannelChatID),
		AdminChatID:       ptrPgInt8(channel.AdminChatID),
		DiscussionsChatID: ptrPgInt8(channel.DiscussionsChatID),
		Decorations:       ptrPgText(channel.Decorations),
	})
	if err != nil {
		return nil, fmt.Errorf("%s:%v", op, err)
	}

	return &entities.Channel{
		ID:                int(updatedChannel.ID),
		Code:              updatedChannel.Code,
		ChannelChatID:     &updatedChannel.ChannelChatID.Int64,
		AdminChatID:       &updatedChannel.AdminChatID.Int64,
		DiscussionsChatID: &updatedChannel.DiscussionsChatID.Int64,
		Decorations:       &updatedChannel.Decorations.String,
	}, nil
}
