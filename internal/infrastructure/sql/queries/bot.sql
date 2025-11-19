-- name: CreateIfNotExists :one
INSERT INTO bots (tgid, type)
VALUES ($1, $2)
ON CONFLICT (tgid, type) DO UPDATE
  SET tgid = EXCLUDED.tgid
RETURNING id, tgid, type, channel_id;
