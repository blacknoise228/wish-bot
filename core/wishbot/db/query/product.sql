-- name: GetProductsByCategory :many
SELECT p.name,
 p.id,
  p.price,
  p.description,
   p.status,
   p.image,
   p.category_id,
   p.created_at,
   p.updated_at,
   p.shop_id,
    p.admin_id,
    s.status_name FROM product as p
LEFT JOIN dim_product_status s ON p.status = s.id
WHERE p.category_id = $1;

-- name: GetProductByID :one
SELECT * FROM product
WHERE id = $1;