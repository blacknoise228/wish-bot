-- name: GetProductsByCategory :many
SELECT * FROM product
WHERE category_id = $1;

-- name: GetProductByID :one
SELECT * FROM product
WHERE id = $1;