package service

import (
	"context"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"
)

type Services struct {
	User User
}

func NewServices(db *db.Queries,
	config *config.Config,
) *Services {
	users := NewUsers(db, config)
	return &Services{
		User: users,
	}
}

type User interface {
	CreateUser(ctx context.Context,
		req db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context,
		chatID int32) (db.User, error)
	UpdateUser(ctx context.Context,
		req db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context,
		chatID int32) error
}
