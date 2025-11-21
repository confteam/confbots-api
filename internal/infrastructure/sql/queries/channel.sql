-- name: CreateChannel :one
WITH create_channel AS (
  INSERT INTO channels (code, channel_chat_id, admin_chat_id, discussions_chat_id)
  VALUES ($1, $2, $3, $4)
  RETURNING *
)
UPDATE bots
SET channel_id = create_channel.id
FROM create_channel
WHERE tgid = $5 AND type = $6
RETURNING create_channel.*;

-- name: UpdateChannel :one
UPDATE channels
SET
  channel_chat_id = COALESCE($2, channel_chat_id),
  admin_chat_id = COALESCE($3, admin_chat_id),
  discussions_chat_id = COALESCE($4, discussions_chat_id),
  decorations = COALESCE($5, decorations)
WHERE id = $1
RETURNING id, code, channel_chat_id, admin_chat_id, discussions_chat_id, decorations;
