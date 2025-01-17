-- name: CreateFriendship :one
INSERT INTO friends ( 
    chat_id,
    friend_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetFriendship :one
SELECT * FROM friends
WHERE chat_id = $1 AND friend_id = $2 LIMIT 1;

-- name: GetAprovedFriendships :many
SELECT 
    u_friend.username AS username,
    u_friend.chat_id AS friend_id
FROM friends f
JOIN users u_friend ON u_friend.chat_id = f.friend_id
JOIN dim_friend_status d ON d.id = f.status
WHERE 
    f.chat_id = $1
    AND f.status = 1

UNION

SELECT 
    u_chat.username AS username,
    u_chat.chat_id AS friend_id
FROM friends f
JOIN users u_chat ON u_chat.chat_id = f.chat_id
JOIN dim_friend_status d ON d.id = f.status
WHERE 
    f.friend_id = $1
    AND f.status = 1;


-- name: GetPendingFriendships :many
SELECT 
    u_friend.username AS username,
    u_friend.chat_id AS friend_id,
    d.status_name
FROM friends f
JOIN users u_friend ON u_friend.chat_id = f.friend_id
JOIN dim_friend_status d ON d.id = f.status
WHERE 
    f.chat_id = $1
    AND f.status != 1

UNION

SELECT 
    u_chat.username AS username,
    u_chat.chat_id AS friend_id,
    d.status_name
FROM friends f
JOIN users u_chat ON u_chat.chat_id = f.chat_id
JOIN dim_friend_status d ON d.id = f.status
WHERE 
    f.friend_id = $1
    AND f.status != 1;


-- name: DeleteFriendship :exec
DELETE FROM friends
WHERE chat_id = $1 AND friend_id = $2;

-- name: UpdateFriendshipStatus :one
UPDATE friends
SET status = $1
WHERE chat_id = $2 AND friend_id = $3
RETURNING *;