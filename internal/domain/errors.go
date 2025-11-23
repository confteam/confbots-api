package domain

import "errors"

var (
	ErrBotNotFound         = errors.New("bot not found")
	ErrUserNotFound        = errors.New("user not found")
	ErrUserChannelNotFound = errors.New("user-channel not found")
	ErrChannelNotFound     = errors.New("channel not found")
	ErrReplyNotFound       = errors.New("reply not found")
	ErrTakeNotFound        = errors.New("take not found")
)
