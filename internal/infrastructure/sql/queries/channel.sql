-- name: CreateChannel :one
INSERT INTO channels (code, channel_chat_id, admin_chat_id, discussions_chat_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateChannel :exec
UPDATE channels
SET
  channel_chat_id = COALESCE($2, channel_chat_id),
  admin_chat_id = COALESCE($3, admin_chat_id),
  discussions_chat_id = COALESCE($4, discussions_chat_id),
  decorations = COALESCE($5, decorations)
WHERE id = $1
RETURNING id, code, channel_chat_id, admin_chat_id, discussions_chat_id, decorations;
