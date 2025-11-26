-- name: CreateReply :one
INSERT INTO replies (user_message_id, admin_message_id, take_id, channel_id)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetReplyByMsgIdAndChannelId :one
SELECT *
FROM replies
WHERE (user_message_id = $1 OR admin_message_id = $1) AND channel_id = $2;

-- name: GetReplyByMsgIdAndTakeId :one
SELECT *
FROM replies
WHERE (user_message_id = $1 OR admin_message_id = $1) AND take_id = $2;
