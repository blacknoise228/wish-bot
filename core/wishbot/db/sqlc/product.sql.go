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

const getProductByID = `-- name: GetProductByID :one
SELECT id, name, price, description, image, category_id, status, shop_id, admin_id, created_at, updated_at FROM product
WHERE id = $1
`

func (q *Queries) GetProductByID(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRow(ctx, getProductByID, id)
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

const getProductsByCategory = `-- name: GetProductsByCategory :many
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
    s.status_name FROM product p
LEFT JOIN dim_product_status s ON p.status = s.id
WHERE p.category_id = $1
`

type GetProductsByCategoryRow struct {
	Name        string           `json:"name"`
	ID          uuid.UUID        `json:"id"`
	Price       float64          `json:"price"`
	Description string           `json:"description"`
	Status      int32            `json:"status"`
	Image       string           `json:"image"`
	CategoryID  int32            `json:"category_id"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	ShopID      uuid.UUID        `json:"shop_id"`
	AdminID     int64            `json:"admin_id"`
	StatusName  pgtype.Text      `json:"status_name"`
}

func (q *Queries) GetProductsByCategory(ctx context.Context, categoryID int32) ([]GetProductsByCategoryRow, error) {
	rows, err := q.db.Query(ctx, getProductsByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsByCategoryRow{}
	for rows.Next() {
		var i GetProductsByCategoryRow
		if err := rows.Scan(
			&i.Name,
			&i.ID,
			&i.Price,
			&i.Description,
			&i.Status,
			&i.Image,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ShopID,
			&i.AdminID,
			&i.StatusName,
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
