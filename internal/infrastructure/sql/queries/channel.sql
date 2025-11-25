-- name: CreateChannel :one
INSERT INTO channels (code, channel_chat_id, admin_chat_id, discussions_chat_id)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: UpdateChannel :one
UPDATE channels
SET
  channel_chat_id = COALESCE($2, channel_chat_id),
  admin_chat_id = COALESCE($3, admin_chat_id),
  discussions_chat_id = COALESCE($4, discussions_chat_id),
  decorations = COALESCE($5, decorations)
WHERE id = $1
RETURNING id, code, channel_chat_id, admin_chat_id, discussions_chat_id, decorations;

-- name: FindChannelByCode :one
SELECT id, channel_chat_id, admin_chat_id, discussions_chat_id, decorations
FROM channels
WHERE code = $1;

-- name: FindChannelById :one
SELECT code, channel_chat_id, admin_chat_id, discussions_chat_id, decorations
FROM channels
WHERE id = $1;

-- name: FindChannelByChatId :one
SELECT id
FROM channels
WHERE admin_chat_id = $1 OR channel_chat_id = $1 OR discussions_chat_id = $1;

-- name: GetAllUserChannels :many
SELECT
    uc.channel_id,
    c.channel_chat_id
FROM user_channels AS uc
JOIN channels AS c ON c.id = uc.channel_id
WHERE uc.user_id = $1;
