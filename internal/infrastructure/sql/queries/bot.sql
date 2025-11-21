-- name: FindBotByTgIdAndType :one
SELECT 
    b.id, b.tgid, b.type, b.channel_id,
    c.id AS channel_id,
    c.code, 
    c.channel_chat_id, 
    c.admin_chat_id, 
    c.discussions_chat_id, 
    c.decorations
FROM bots b
LEFT JOIN channels c ON b.channel_id = c.id
WHERE b.tgid = $1 AND b.type = $2;

-- name: CreateBot :one
INSERT INTO bots (tgid, type)
VALUES ($1, $2)
RETURNING id, tgid, type, channel_id;

-- name: UpdateBotChannelID :exec
UPDATE bots
SET channel_id = $1
WHERE tgid = $2 AND type = $3;
