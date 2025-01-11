// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: shop.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createShop = `-- name: CreateShop :one
INSERT INTO shop ( 
    name,
    description,
    image,
    token
) VALUES (
    $1, $2, $3, $4
) RETURNING id, name, description, image, token, created_at, updated_at
`

type CreateShopParams struct {
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
	Image       string      `json:"image"`
	Token       string      `json:"token"`
}

func (q *Queries) CreateShop(ctx context.Context, arg CreateShopParams) (Shop, error) {
	row := q.db.QueryRow(ctx, createShop,
		arg.Name,
		arg.Description,
		arg.Image,
		arg.Token,
	)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Token,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createShopAdmin = `-- name: CreateShopAdmin :exec
INSERT INTO shop_admins (
    admin_id,
    shop_id
) VALUES (
    $1, $2
)
`

type CreateShopAdminParams struct {
	AdminID int64     `json:"admin_id"`
	ShopID  uuid.UUID `json:"shop_id"`
}

func (q *Queries) CreateShopAdmin(ctx context.Context, arg CreateShopAdminParams) error {
	_, err := q.db.Exec(ctx, createShopAdmin, arg.AdminID, arg.ShopID)
	return err
}

const deleteShop = `-- name: DeleteShop :exec
DELETE FROM shop
WHERE id = $1 AND token = $2
`

type DeleteShopParams struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

func (q *Queries) DeleteShop(ctx context.Context, arg DeleteShopParams) error {
	_, err := q.db.Exec(ctx, deleteShop, arg.ID, arg.Token)
	return err
}

const getRandomAdminByShopID = `-- name: GetRandomAdminByShopID :one
SELECT admin_id, shop_id FROM shop_admins
WHERE shop_id = $1
ORDER BY random()
LIMIT 1
`

func (q *Queries) GetRandomAdminByShopID(ctx context.Context, shopID uuid.UUID) (ShopAdmin, error) {
	row := q.db.QueryRow(ctx, getRandomAdminByShopID, shopID)
	var i ShopAdmin
	err := row.Scan(&i.AdminID, &i.ShopID)
	return i, err
}

const getShopAdminsByShopID = `-- name: GetShopAdminsByShopID :many
SELECT admin_id, shop_id FROM shop_admins
WHERE shop_id = $1
`

func (q *Queries) GetShopAdminsByShopID(ctx context.Context, shopID uuid.UUID) ([]ShopAdmin, error) {
	rows, err := q.db.Query(ctx, getShopAdminsByShopID, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ShopAdmin{}
	for rows.Next() {
		var i ShopAdmin
		if err := rows.Scan(&i.AdminID, &i.ShopID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getShopByID = `-- name: GetShopByID :one
SELECT id, name, description, image, token, created_at, updated_at FROM shop
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetShopByID(ctx context.Context, id uuid.UUID) (Shop, error) {
	row := q.db.QueryRow(ctx, getShopByID, id)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Token,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getShopByToken = `-- name: GetShopByToken :one
SELECT id, name, description, image, token, created_at, updated_at FROM shop
WHERE token = $1 LIMIT 1
`

func (q *Queries) GetShopByToken(ctx context.Context, token string) (Shop, error) {
	row := q.db.QueryRow(ctx, getShopByToken, token)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Token,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getShops = `-- name: GetShops :many
SELECT id, name, description, image, token, created_at, updated_at FROM shop
`

func (q *Queries) GetShops(ctx context.Context) ([]Shop, error) {
	rows, err := q.db.Query(ctx, getShops)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Shop{}
	for rows.Next() {
		var i Shop
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Image,
			&i.Token,
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

const updateShop = `-- name: UpdateShop :one
UPDATE shop
SET
name = $2,
image = $3,
updated_at = now()
WHERE id = $1
RETURNING id, name, description, image, token, created_at, updated_at
`

type UpdateShopParams struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}

func (q *Queries) UpdateShop(ctx context.Context, arg UpdateShopParams) (Shop, error) {
	row := q.db.QueryRow(ctx, updateShop, arg.ID, arg.Name, arg.Image)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Token,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
