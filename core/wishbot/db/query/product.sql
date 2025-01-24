-- name: GetProductsByCategory :many
SELECT p.*, s.status_name FROM product p
LEFT JOIN dim_product_status s ON p.status = s.id
WHERE p.category_id = $1;

-- name: GetProductByID :one
SELECT * FROM product
WHERE id = $1;