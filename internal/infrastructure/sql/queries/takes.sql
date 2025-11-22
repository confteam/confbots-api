-- name: CreateTake :one
INSERT INTO takes (user_message_id, admin_message_id, user_channel_id, channel_id)
VALUES ($1, $2, $3, $4)
RETURNING id, user_message_id, admin_message_id, user_channel_id, channel_id;

-- name: GetTakeById :one
SELECT id, status, user_message_id, admin_message_id, user_channel_id, channel_id
FROM takes
WHERE id = $1 AND channel_id = $2;

-- name: GetTakeByMsgId :one
SELECT id, status, user_message_id, admin_message_id, user_channel_id, channel_id
FROM takes
WHERE (user_message_id = $1 OR admin_message_id = $1) AND channel_id = $2;

-- name: UpdateTakeStatus :exec
UPDATE takes
SET status = $1
WHERE id = $2 AND channel_id = $3;
