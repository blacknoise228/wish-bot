package service

import (
	"context"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"
)

type Users struct {
	db     *db.Queries
	config *config.Config
}

func NewUsers(db *db.Queries, config *config.Config) User {
	return &Users{
		db:     db,
		config: config,
	}
}

func (u *Users) CreateUser(ctx context.Context,
	req db.CreateUserParams) (db.User, error) {
	return u.db.CreateUser(ctx, req)
}

func (u *Users) GetUser(ctx context.Context,
	chatID int32) (db.User, error) {
	return u.db.GetUser(ctx, chatID)
}

func (u *Users) UpdateUser(ctx context.Context,
	req db.UpdateUserParams) (db.User, error) {
	return u.db.UpdateUser(ctx, req)
}

func (u *Users) DeleteUser(ctx context.Context,
	chatID int32) error {
	return u.db.DeleteUser(ctx, chatID)
}
