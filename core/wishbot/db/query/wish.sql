-- name: CreateWish :one
INSERT INTO wish (
    chat_id, 
    product_id,
    status
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetWishesForUser :many
SELECT 
    w.chat_id,
    w.product_id, 
    d.status_name, 
    w.id, 
    w.created_at, 
    u.username 
FROM wish w
JOIN users u ON w.chat_id = u.chat_id
JOIN dim_wish_status d ON w.status = d.id
WHERE w.chat_id = $1;


-- name: GetWishesPublic :many
SELECT w.product_id, w.id, d.status_name, w.created_at, u.username
FROM wish w
JOIN users u ON w.chat_id = u.chat_id
JOIN dim_wish_status d ON w.status = d.id
WHERE w.chat_id = $1
  AND (
      w.status = 1 
      OR (
          EXISTS (
              SELECT 1
              FROM friends f
              WHERE (
                  (f.chat_id = $2 AND f.friend_id = w.chat_id)
                  OR (f.chat_id = w.chat_id AND f.friend_id = $2)
              )
              AND f.status = 1 
          )
      )
  );


-- name: DeleteWish :exec
DELETE FROM wish
WHERE chat_id = $1 AND id = $2;

-- name: UpdateWishStatus :one
UPDATE wish
SET 
status = $1
WHERE chat_id = $2 AND id = $3
RETURNING *;

-- name: GetWish :one
SELECT * FROM wish
WHERE chat_id = $1 AND id = $2;

-- name: GetWishByID :one
SELECT * FROM wish
WHERE id = $1;