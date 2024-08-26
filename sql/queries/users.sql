-- name: InsertUser :one
INSERT INTO users
  (id, email)
VALUES
  ($1, $2)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUserEmailByID :exec
UPDATE users
SET
  email = $2
WHERE id = $1;

-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;