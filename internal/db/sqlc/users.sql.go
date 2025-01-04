// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, 
    chat_id
) VALUES (
    $1, $2
) RETURNING username, chat_id, created_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	ChatID   int32  `json:"chat_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.ChatID)
	var i User
	err := row.Scan(&i.Username, &i.ChatID, &i.CreatedAt)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE chat_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, chatID int32) error {
	_, err := q.db.Exec(ctx, deleteUser, chatID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT username, chat_id, created_at FROM users
WHERE chat_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, chatID int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, chatID)
	var i User
	err := row.Scan(&i.Username, &i.ChatID, &i.CreatedAt)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $1
WHERE chat_id = $2
RETURNING username, chat_id, created_at
`

type UpdateUserParams struct {
	Username string `json:"username"`
	ChatID   int32  `json:"chat_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.Username, arg.ChatID)
	var i User
	err := row.Scan(&i.Username, &i.ChatID, &i.CreatedAt)
	return i, err
}
