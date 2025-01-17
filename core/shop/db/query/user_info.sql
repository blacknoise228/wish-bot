-- name: GetUserInfo :one
SELECT * FROM user_info
WHERE chat_id = $1 LIMIT 1;