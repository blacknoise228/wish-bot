package service

import (
	"context"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"
)

type Services struct {
	User   User
	Wish   Wisher
	Friend Friender
}

func NewServices(db *db.Queries,
	config *config.Config,
) *Services {
	return &Services{
		User:   NewUsers(db, config),
		Wish:   NewWish(db, config),
		Friend: NewFriend(db, config),
	}
}

type User interface {
	CreateUser(ctx context.Context,
		req db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context,
		chatID int64) (db.User, error)
	GetUserByUsername(ctx context.Context,
		username string) (*db.User, error)
	UpdateUser(ctx context.Context,
		req db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context,
		chatID int64) error
}

type Wisher interface {
	CreateWish(ctx context.Context,
		req db.CreateWishParams) (db.Wish, error)
	GetWishesForUser(ctx context.Context,
		chatID int64) ([]db.GetWishesForUserRow, error)
	GetUserWishes(ctx context.Context,
		chatID int64,
		friendID int64,
	) []db.GetWishesPublicRow
	DeleteWish(ctx context.Context,
		chatID int64,
		wishID int) error
}

type Friender interface {
	CreateFriendship(ctx context.Context,
		req db.CreateFriendshipParams) (db.Friend, error)
	GetFriendship(ctx context.Context,
		chatID int64,
		friendID int64,
	) (db.Friend, error)
	GetAprovedFriendships(ctx context.Context,
		chatID int64) ([]db.GetAprovedFriendshipsRow, error)
	GetPendingFriendships(ctx context.Context,
		chatID int64) ([]db.GetPendingFriendshipsRow, error)
	DeleteFriendship(ctx context.Context,
		chatID int64,
		friendID int64,
	) error
	UpdateFriendshipStatus(ctx context.Context,
		req db.UpdateFriendshipStatusParams) (db.Friend, error)
}
