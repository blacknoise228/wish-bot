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
status = $3,
updated_at = now()
WHERE customer_id = $2 and id = $1
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrdersByCustomer :many
SELECT * FROM orders
WHERE customer_id = $1;

-- name: GetRandomAdminByShopID :one
SELECT * FROM shop_admins
WHERE shop_id = $1
ORDER BY random()
LIMIT 1;

-- name: GetDimOrderStatusByID :one
SELECT * FROM dim_order_status
WHERE id = $1 LIMIT 1;