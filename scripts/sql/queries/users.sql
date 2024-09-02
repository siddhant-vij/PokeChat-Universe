-- name: InsertUser :exec
INSERT INTO users
  (id, auth_id, username, email, picture_url)
VALUES
  ($1, $2, $3, $4, $5)
ON CONFLICT (auth_id) DO NOTHING;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserIdFromAuthID :one
SELECT id FROM users
WHERE auth_id = $1 LIMIT 1;

-- name: UpdateUserEmailByID :exec
UPDATE users
SET
  email = $2,
  updated_at = NOW()
WHERE id = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;