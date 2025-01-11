-- name: CreateOrder :one
INSERT INTO orders (
    price,
    status,
    customer_id,
    customer_login,
    consignee_id,
    product_id,
    admin_id,
    shop_id
) VALUES (
$1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET 
status = $1,
updated_at = now()
WHERE id = $2 AND admin_id = $3
RETURNING *;

-- name: GetOrdersByShop :many
SELECT * FROM orders
WHERE shop_id = $1;

-- name: GetOrdersByAdmin :many
SELECT * FROM orders
WHERE admin_id = $1;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;