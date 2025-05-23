// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: order.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createOrder = `-- name: CreateOrder :one
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
) RETURNING id, price, status, customer_id, customer_login, consignee_id, product_id, admin_id, shop_id, created_at, updated_at
`

type CreateOrderParams struct {
	Price         float64   `json:"price"`
	Status        int32     `json:"status"`
	CustomerID    int64     `json:"customer_id"`
	CustomerLogin string    `json:"customer_login"`
	ConsigneeID   int64     `json:"consignee_id"`
	ProductID     uuid.UUID `json:"product_id"`
	AdminID       int64     `json:"admin_id"`
	ShopID        uuid.UUID `json:"shop_id"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder,
		arg.Price,
		arg.Status,
		arg.CustomerID,
		arg.CustomerLogin,
		arg.ConsigneeID,
		arg.ProductID,
		arg.AdminID,
		arg.ShopID,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Price,
		&i.Status,
		&i.CustomerID,
		&i.CustomerLogin,
		&i.ConsigneeID,
		&i.ProductID,
		&i.AdminID,
		&i.ShopID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDimOrderStatusByID = `-- name: GetDimOrderStatusByID :one
SELECT id, status_code, status_name FROM dim_order_status
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDimOrderStatusByID(ctx context.Context, id int32) (DimOrderStatus, error) {
	row := q.db.QueryRow(ctx, getDimOrderStatusByID, id)
	var i DimOrderStatus
	err := row.Scan(&i.ID, &i.StatusCode, &i.StatusName)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT id, price, status, customer_id, customer_login, consignee_id, product_id, admin_id, shop_id, created_at, updated_at FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Price,
		&i.Status,
		&i.CustomerID,
		&i.CustomerLogin,
		&i.ConsigneeID,
		&i.ProductID,
		&i.AdminID,
		&i.ShopID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getOrdersByAdmin = `-- name: GetOrdersByAdmin :many
SELECT o.id, price, status, customer_id, customer_login, consignee_id, product_id, admin_id, shop_id, created_at, updated_at, dos.id, status_code, status_name FROM orders o
LEFT JOIN dim_order_status dos ON o.status = dos.id
WHERE o.admin_id = $1
`

type GetOrdersByAdminRow struct {
	ID            uuid.UUID        `json:"id"`
	Price         float64          `json:"price"`
	Status        int32            `json:"status"`
	CustomerID    int64            `json:"customer_id"`
	CustomerLogin string           `json:"customer_login"`
	ConsigneeID   int64            `json:"consignee_id"`
	ProductID     uuid.UUID        `json:"product_id"`
	AdminID       int64            `json:"admin_id"`
	ShopID        uuid.UUID        `json:"shop_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	ID_2          pgtype.Int4      `json:"id_2"`
	StatusCode    pgtype.Text      `json:"status_code"`
	StatusName    pgtype.Text      `json:"status_name"`
}

func (q *Queries) GetOrdersByAdmin(ctx context.Context, adminID int64) ([]GetOrdersByAdminRow, error) {
	rows, err := q.db.Query(ctx, getOrdersByAdmin, adminID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrdersByAdminRow{}
	for rows.Next() {
		var i GetOrdersByAdminRow
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Status,
			&i.CustomerID,
			&i.CustomerLogin,
			&i.ConsigneeID,
			&i.ProductID,
			&i.AdminID,
			&i.ShopID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.StatusCode,
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

const getOrdersByShop = `-- name: GetOrdersByShop :many
SELECT o.id, price, status, customer_id, customer_login, consignee_id, product_id, admin_id, shop_id, created_at, updated_at, dos.id, status_code, status_name FROM orders o
LEFT JOIN dim_order_status dos ON o.status = dos.id
WHERE o.shop_id = $1
`

type GetOrdersByShopRow struct {
	ID            uuid.UUID        `json:"id"`
	Price         float64          `json:"price"`
	Status        int32            `json:"status"`
	CustomerID    int64            `json:"customer_id"`
	CustomerLogin string           `json:"customer_login"`
	ConsigneeID   int64            `json:"consignee_id"`
	ProductID     uuid.UUID        `json:"product_id"`
	AdminID       int64            `json:"admin_id"`
	ShopID        uuid.UUID        `json:"shop_id"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
	ID_2          pgtype.Int4      `json:"id_2"`
	StatusCode    pgtype.Text      `json:"status_code"`
	StatusName    pgtype.Text      `json:"status_name"`
}

func (q *Queries) GetOrdersByShop(ctx context.Context, shopID uuid.UUID) ([]GetOrdersByShopRow, error) {
	rows, err := q.db.Query(ctx, getOrdersByShop, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrdersByShopRow{}
	for rows.Next() {
		var i GetOrdersByShopRow
		if err := rows.Scan(
			&i.ID,
			&i.Price,
			&i.Status,
			&i.CustomerID,
			&i.CustomerLogin,
			&i.ConsigneeID,
			&i.ProductID,
			&i.AdminID,
			&i.ShopID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.StatusCode,
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

const updateOrderStatus = `-- name: UpdateOrderStatus :exec
UPDATE orders
SET 
status = $1,
updated_at = now()
WHERE id = $2 AND admin_id = $3
RETURNING id, price, status, customer_id, customer_login, consignee_id, product_id, admin_id, shop_id, created_at, updated_at
`

type UpdateOrderStatusParams struct {
	Status  int32     `json:"status"`
	ID      uuid.UUID `json:"id"`
	AdminID int64     `json:"admin_id"`
}

func (q *Queries) UpdateOrderStatus(ctx context.Context, arg UpdateOrderStatusParams) error {
	_, err := q.db.Exec(ctx, updateOrderStatus, arg.Status, arg.ID, arg.AdminID)
	return err
}
