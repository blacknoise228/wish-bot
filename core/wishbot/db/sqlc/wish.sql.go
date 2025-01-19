// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: wish.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createWish = `-- name: CreateWish :one
INSERT INTO wish (
    chat_id, 
    product_id,
    status
) VALUES (
    $1, $2, $3
) RETURNING id, chat_id, created_at, product_id, status
`

type CreateWishParams struct {
	ChatID    int64     `json:"chat_id"`
	ProductID uuid.UUID `json:"product_id"`
	Status    int32     `json:"status"`
}

func (q *Queries) CreateWish(ctx context.Context, arg CreateWishParams) (Wish, error) {
	row := q.db.QueryRow(ctx, createWish, arg.ChatID, arg.ProductID, arg.Status)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.ProductID,
		&i.Status,
	)
	return i, err
}

const deleteWish = `-- name: DeleteWish :exec
DELETE FROM wish
WHERE chat_id = $1 AND id = $2
`

type DeleteWishParams struct {
	ChatID int64 `json:"chat_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) DeleteWish(ctx context.Context, arg DeleteWishParams) error {
	_, err := q.db.Exec(ctx, deleteWish, arg.ChatID, arg.ID)
	return err
}

const getWish = `-- name: GetWish :one
SELECT id, chat_id, created_at, product_id, status FROM wish
WHERE chat_id = $1 AND id = $2
`

type GetWishParams struct {
	ChatID int64 `json:"chat_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) GetWish(ctx context.Context, arg GetWishParams) (Wish, error) {
	row := q.db.QueryRow(ctx, getWish, arg.ChatID, arg.ID)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.ProductID,
		&i.Status,
	)
	return i, err
}

const getWishByID = `-- name: GetWishByID :one
SELECT id, chat_id, created_at, product_id, status FROM wish
WHERE id = $1
`

func (q *Queries) GetWishByID(ctx context.Context, id int32) (Wish, error) {
	row := q.db.QueryRow(ctx, getWishByID, id)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.ProductID,
		&i.Status,
	)
	return i, err
}

const getWishesForUser = `-- name: GetWishesForUser :many
SELECT 
    w.chat_id,
    w.product_id, 
    d.status_name, 
    w.id, 
    w.created_at, 
    u.username 
FROM wish w
JOIN users u ON w.chat_id = u.chat_id
JOIN dim_wish_status d ON w.status = d.id
WHERE w.chat_id = $1
`

type GetWishesForUserRow struct {
	ChatID     int64     `json:"chat_id"`
	ProductID  uuid.UUID `json:"product_id"`
	StatusName string    `json:"status_name"`
	ID         int32     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Username   string    `json:"username"`
}

func (q *Queries) GetWishesForUser(ctx context.Context, chatID int64) ([]GetWishesForUserRow, error) {
	rows, err := q.db.Query(ctx, getWishesForUser, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetWishesForUserRow{}
	for rows.Next() {
		var i GetWishesForUserRow
		if err := rows.Scan(
			&i.ChatID,
			&i.ProductID,
			&i.StatusName,
			&i.ID,
			&i.CreatedAt,
			&i.Username,
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

const getWishesPublic = `-- name: GetWishesPublic :many
SELECT w.product_id, w.id, d.status_name, w.created_at, u.username
FROM wish w
JOIN users u ON w.chat_id = u.chat_id
JOIN dim_wish_status d ON w.status = d.id
WHERE w.chat_id = $1
  AND (
      w.status = 1 
      OR (
          EXISTS (
              SELECT 1
              FROM friends f
              WHERE (
                  (f.chat_id = $2 AND f.friend_id = w.chat_id)
                  OR (f.chat_id = w.chat_id AND f.friend_id = $2)
              )
              AND f.status = 1 
          )
      )
  )
`

type GetWishesPublicParams struct {
	ChatID   int64 `json:"chat_id"`
	ChatID_2 int64 `json:"chat_id_2"`
}

type GetWishesPublicRow struct {
	ProductID  uuid.UUID `json:"product_id"`
	ID         int32     `json:"id"`
	StatusName string    `json:"status_name"`
	CreatedAt  time.Time `json:"created_at"`
	Username   string    `json:"username"`
}

func (q *Queries) GetWishesPublic(ctx context.Context, arg GetWishesPublicParams) ([]GetWishesPublicRow, error) {
	rows, err := q.db.Query(ctx, getWishesPublic, arg.ChatID, arg.ChatID_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetWishesPublicRow{}
	for rows.Next() {
		var i GetWishesPublicRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ID,
			&i.StatusName,
			&i.CreatedAt,
			&i.Username,
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

const updateWishStatus = `-- name: UpdateWishStatus :one
UPDATE wish
SET 
status = $1
WHERE chat_id = $2 AND id = $3
RETURNING id, chat_id, created_at, product_id, status
`

type UpdateWishStatusParams struct {
	Status int32 `json:"status"`
	ChatID int64 `json:"chat_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateWishStatus(ctx context.Context, arg UpdateWishStatusParams) (Wish, error) {
	row := q.db.QueryRow(ctx, updateWishStatus, arg.Status, arg.ChatID, arg.ID)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.ProductID,
		&i.Status,
	)
	return i, err
}
