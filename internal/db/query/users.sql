-- name: CreateUser :one
INSERT INTO users (
    username, 
    chat_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE chat_id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET username = $1
WHERE chat_id = $2
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE chat_id = $1;