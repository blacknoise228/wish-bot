package service

import (
	"context"
	"errors"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"

	"github.com/jackc/pgx/v5"
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
	chatID int64) (db.User, error) {
	return u.db.GetUser(ctx, chatID)
}

func (u *Users) GetUserByUsername(ctx context.Context,
	username string) (*db.User, error) {
	user, err := u.db.GetUserByUsername(ctx, username)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (u *Users) UpdateUser(ctx context.Context,
	req db.UpdateUserParams) (db.User, error) {
	return u.db.UpdateUser(ctx, req)
}

func (u *Users) DeleteUser(ctx context.Context,
	chatID int64) error {
	return u.db.DeleteUser(ctx, chatID)
}
