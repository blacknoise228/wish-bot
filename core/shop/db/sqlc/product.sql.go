// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: product.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
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
) RETURNING id, name, price, description, image, category_id, status, shop_id, admin_id, created_at, updated_at
`

type CreateProductParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	CategoryID  int32     `json:"category_id"`
	Status      int32     `json:"status"`
	AdminID     int64     `json:"admin_id"`
	ShopID      uuid.UUID `json:"shop_id"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Image,
		arg.CategoryID,
		arg.Status,
		arg.AdminID,
		arg.ShopID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CategoryID,
		&i.Status,
		&i.ShopID,
		&i.AdminID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1 AND shop_id = $2
`

type DeleteProductParams struct {
	ID     uuid.UUID `json:"id"`
	ShopID uuid.UUID `json:"shop_id"`
}

func (q *Queries) DeleteProduct(ctx context.Context, arg DeleteProductParams) error {
	_, err := q.db.Exec(ctx, deleteProduct, arg.ID, arg.ShopID)
	return err
}

const getProductByID = `-- name: GetProductByID :one
SELECT p.id, name, price, description, image, category_id, status, shop_id, admin_id, created_at, updated_at, dps.id, status_code, status_name FROM product p
LEFT JOIN dim_product_status dps ON p.status = dps.id
WHERE p.id = $1
`

type GetProductByIDRow struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Price       float64          `json:"price"`
	Description string           `json:"description"`
	Image       string           `json:"image"`
	CategoryID  int32            `json:"category_id"`
	Status      int32            `json:"status"`
	ShopID      uuid.UUID        `json:"shop_id"`
	AdminID     int64            `json:"admin_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	ID_2        pgtype.Int4      `json:"id_2"`
	StatusCode  pgtype.Text      `json:"status_code"`
	StatusName  pgtype.Text      `json:"status_name"`
}

func (q *Queries) GetProductByID(ctx context.Context, id uuid.UUID) (GetProductByIDRow, error) {
	row := q.db.QueryRow(ctx, getProductByID, id)
	var i GetProductByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CategoryID,
		&i.Status,
		&i.ShopID,
		&i.AdminID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.StatusCode,
		&i.StatusName,
	)
	return i, err
}

const getProductCategoryByName = `-- name: GetProductCategoryByName :one
SELECT id, category_code, category_name FROM dim_product_category
WHERE category_name = $1 LIMIT 1
`

func (q *Queries) GetProductCategoryByName(ctx context.Context, categoryName string) (DimProductCategory, error) {
	row := q.db.QueryRow(ctx, getProductCategoryByName, categoryName)
	var i DimProductCategory
	err := row.Scan(&i.ID, &i.CategoryCode, &i.CategoryName)
	return i, err
}

const getProductStatusByName = `-- name: GetProductStatusByName :one
SELECT id, status_code, status_name FROM dim_product_status
WHERE status_name = $1 LIMIT 1
`

func (q *Queries) GetProductStatusByName(ctx context.Context, statusName string) (DimProductStatus, error) {
	row := q.db.QueryRow(ctx, getProductStatusByName, statusName)
	var i DimProductStatus
	err := row.Scan(&i.ID, &i.StatusCode, &i.StatusName)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, price, description, image, category_id, status, shop_id, admin_id, created_at, updated_at FROM product
WHERE shop_id = $1
`

func (q *Queries) GetProducts(ctx context.Context, shopID uuid.UUID) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProducts, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			&i.Description,
			&i.Image,
			&i.CategoryID,
			&i.Status,
			&i.ShopID,
			&i.AdminID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
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
RETURNING id, name, price, description, image, category_id, status, shop_id, admin_id, created_at, updated_at
`

type UpdateProductParams struct {
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	CategoryID  int32     `json:"category_id"`
	Status      int32     `json:"status"`
	AdminID     int64     `json:"admin_id"`
	ID          uuid.UUID `json:"id"`
	ShopID      uuid.UUID `json:"shop_id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.Name,
		arg.Price,
		arg.Image,
		arg.Description,
		arg.CategoryID,
		arg.Status,
		arg.AdminID,
		arg.ID,
		arg.ShopID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CategoryID,
		&i.Status,
		&i.ShopID,
		&i.AdminID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProductStatus = `-- name: UpdateProductStatus :exec
UPDATE product
SET 
status = $1
WHERE id = $2 AND shop_id = $3
RETURNING id, name, price, description, image, category_id, status, shop_id, admin_id, created_at, updated_at
`

type UpdateProductStatusParams struct {
	Status int32     `json:"status"`
	ID     uuid.UUID `json:"id"`
	ShopID uuid.UUID `json:"shop_id"`
}

func (q *Queries) UpdateProductStatus(ctx context.Context, arg UpdateProductStatusParams) error {
	_, err := q.db.Exec(ctx, updateProductStatus, arg.Status, arg.ID, arg.ShopID)
	return err
}
