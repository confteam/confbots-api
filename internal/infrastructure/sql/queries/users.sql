-- name: UpsertUser :one
WITH upsert_user AS (
  INSERT INTO users (tgid)
  VALUES ($1)
  ON CONFLICT (tgid) DO UPDATE
  SET tgid = EXCLUDED.tgid
  RETURNING id
),
upsert_user_channel AS (
  INSERT INTO user_channels (user_id, channel_id, role)
  SELECT id, $2, $3 FROM upsert_user
  ON CONFLICT (user_id, channel_id) DO UPDATE
  SET role = EXCLUDED.role
)
SELECT id FROM upsert_user;

-- name: GetUserIdByTgId :one
SELECT id
FROM users
WHERE tgid = $1;

-- name: UpdateUserRole :exec
UPDATE user_channels
SET role = $1
WHERE user_id = $2 AND channel_id = $3;

-- name: GetUserRole :one
SELECT role
FROM user_channels
WHERE user_id = $1 AND channel_id = $2;
