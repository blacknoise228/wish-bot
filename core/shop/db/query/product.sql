-- name: CreateProduct :one
INSERT INTO product (
    name, 
    description,
    price,
    image,
    category_id,
    status,
    admin_id,
    shop_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1 AND shop_id = $2;

-- name: UpdateProductStatus :exec
UPDATE product
SET 
status = $1
WHERE id = $2 AND shop_id = $3
RETURNING *;

-- name: UpdateProduct :one
UPDATE product
SET 
name = $1,
price = $2,
image = $3,
description = $4,
category_id = $5,
status = $6,
admin_id = $7,
updated_at = now()
WHERE id = $8 AND shop_id = $9
RETURNING *;

-- name: GetProducts :many
SELECT * FROM product
WHERE shop_id = $1;

-- name: GetProductCategoryByName :one
SELECT * FROM dim_product_category
WHERE category_name = $1 LIMIT 1;

-- name: GetProductStatusByName :one
SELECT * FROM dim_product_status
WHERE status_name = $1 LIMIT 1;