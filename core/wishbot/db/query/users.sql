-- name: CreateUser :one
INSERT INTO users (
    username, 
    chat_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: CreateUserInfo :exec
INSERT INTO user_info (
    chat_id,
    address,
    phone,
    name,
    description
) VALUES (
    $1, $2, $3, $4, $5
);

-- name: UpdateUserInfoAddress :exec
UPDATE user_info 
SET address = $1
WHERE chat_id = $2;

-- name: UpdateUserInfoPhone :exec
UPDATE user_info 
SET phone = $1
WHERE chat_id = $2;

-- name: UpdateUserInfoName :exec
UPDATE user_info 
SET name = $1
WHERE chat_id = $2;

-- name: UpdateUserInfoDescription :exec
UPDATE user_info 
SET description = $1
WHERE chat_id = $2;

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

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;