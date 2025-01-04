// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: wish.sql

package db

import (
	"context"
	"time"
)

const createWish = `-- name: CreateWish :one
INSERT INTO wish (
    chat_id, 
    description,
    link,
    status
) VALUES (
    $1, $2, $3, $4
) RETURNING id, chat_id, created_at, description, link, status
`

type CreateWishParams struct {
	ChatID      int32  `json:"chat_id"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Status      string `json:"status"`
}

func (q *Queries) CreateWish(ctx context.Context, arg CreateWishParams) (Wish, error) {
	row := q.db.QueryRow(ctx, createWish,
		arg.ChatID,
		arg.Description,
		arg.Link,
		arg.Status,
	)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.Description,
		&i.Link,
		&i.Status,
	)
	return i, err
}

const deleteWish = `-- name: DeleteWish :exec
DELETE FROM wish
WHERE chat_id = $1 AND id = $2
`

type DeleteWishParams struct {
	ChatID int32 `json:"chat_id"`
	ID     int32 `json:"id"`
}

func (q *Queries) DeleteWish(ctx context.Context, arg DeleteWishParams) error {
	_, err := q.db.Exec(ctx, deleteWish, arg.ChatID, arg.ID)
	return err
}

const getWishesForUser = `-- name: GetWishesForUser :many
SELECT id, chat_id, created_at, description, link, status FROM wish
WHERE chat_id = $1
`

func (q *Queries) GetWishesForUser(ctx context.Context, chatID int32) ([]Wish, error) {
	rows, err := q.db.Query(ctx, getWishesForUser, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wish{}
	for rows.Next() {
		var i Wish
		if err := rows.Scan(
			&i.ID,
			&i.ChatID,
			&i.CreatedAt,
			&i.Description,
			&i.Link,
			&i.Status,
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
SELECT w.description, w.link, d.status_name, w.created_at, u.username
FROM wish w
JOIN users u ON w.chat_id = u.chat_id
JOIN dim_wish_status d ON w.status = d.id
WHERE w.chat_id = $1
  AND (
      w.status = 1
      OR EXISTS (
          SELECT 1
          FROM friends f
          WHERE (
              (f.chat_id = $2 AND f.friend_id = w.chat_id)
              OR (f.chat_id = w.chat_id AND f.friend_id = $2)
          )
          AND f.status = 1
      )
  )
`

type GetWishesPublicParams struct {
	ChatID   int32 `json:"chat_id"`
	ChatID_2 int32 `json:"chat_id_2"`
}

type GetWishesPublicRow struct {
	Description string    `json:"description"`
	Link        string    `json:"link"`
	StatusName  string    `json:"status_name"`
	CreatedAt   time.Time `json:"created_at"`
	Username    string    `json:"username"`
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
			&i.Description,
			&i.Link,
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

const updateWish = `-- name: UpdateWish :one
UPDATE wish
SET 
description = $1,
link = $2,
status = $3
WHERE chat_id = $4 AND id = $5
RETURNING id, chat_id, created_at, description, link, status
`

type UpdateWishParams struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	Status      string `json:"status"`
	ChatID      int32  `json:"chat_id"`
	ID          int32  `json:"id"`
}

func (q *Queries) UpdateWish(ctx context.Context, arg UpdateWishParams) (Wish, error) {
	row := q.db.QueryRow(ctx, updateWish,
		arg.Description,
		arg.Link,
		arg.Status,
		arg.ChatID,
		arg.ID,
	)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.Description,
		&i.Link,
		&i.Status,
	)
	return i, err
}

const updateWishStatus = `-- name: UpdateWishStatus :one
UPDATE wish
SET 
status = $1
WHERE chat_id = $2 AND id = $3
RETURNING id, chat_id, created_at, description, link, status
`

type UpdateWishStatusParams struct {
	Status string `json:"status"`
	ChatID int32  `json:"chat_id"`
	ID     int32  `json:"id"`
}

func (q *Queries) UpdateWishStatus(ctx context.Context, arg UpdateWishStatusParams) (Wish, error) {
	row := q.db.QueryRow(ctx, updateWishStatus, arg.Status, arg.ChatID, arg.ID)
	var i Wish
	err := row.Scan(
		&i.ID,
		&i.ChatID,
		&i.CreatedAt,
		&i.Description,
		&i.Link,
		&i.Status,
	)
	return i, err
}
