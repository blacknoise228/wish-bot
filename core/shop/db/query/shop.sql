-- name: CreateShop :one
INSERT INTO shop ( 
    name,
    description,
    image,
    token
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetShopByID :one
SELECT * FROM shop
WHERE id = $1 LIMIT 1;

-- name: GetShopByToken :one
SELECT * FROM shop
WHERE token = $1;

-- name: GetShops :many
SELECT * FROM shop;

-- name: DeleteShop :exec
DELETE FROM shop
WHERE id = $1 AND token = $2;

-- name: UpdateShop :one
UPDATE shop
SET
name = $2,
image = $3,
updated_at = now()
WHERE id = $1
RETURNING *;

-- name: CreateShopAdmin :exec
INSERT INTO shop_admins (
    admin_id,
    shop_id
) VALUES (
    $1, $2
);

-- name: DeleteShopAdmin :exec
DELETE FROM shop_admins
WHERE admin_id = $1;


-- name: GetShopAdminsByShopID :many
SELECT * FROM shop_admins
WHERE shop_id = $1;

-- name: GetShopAdminsByAdminID :one
SELECT * FROM shop_admins
WHERE admin_id = $1
LIMIT 1;

-- name: GetRandomAdminByShopID :one
SELECT * FROM shop_admins
WHERE shop_id = $1
ORDER BY random()
LIMIT 1;