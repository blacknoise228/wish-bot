package service

import (
	"context"
	"wish-bot/core/wishbot/config"
	db "wish-bot/core/wishbot/db/sqlc"
)

type Friend struct {
	db     *db.Queries
	config *config.Config
}

func NewFriend(db *db.Queries, config *config.Config) Friender {
	return &Friend{
		db:     db,
		config: config,
	}
}

func (f *Friend) CreateFriendship(ctx context.Context,
	req db.CreateFriendshipParams) (db.Friend, error) {
	return f.db.CreateFriendship(ctx, req)
}

func (f *Friend) GetFriendship(ctx context.Context,
	chatID int64,
	friendID int64,
) (db.Friend, error) {
	return f.db.GetFriendship(ctx, db.GetFriendshipParams{
		ChatID:   chatID,
		FriendID: friendID,
	})
}

func (f *Friend) DeleteFriendship(ctx context.Context,
	chatID int64,
	friendID int64,
) error {
	return f.db.DeleteFriendship(ctx, db.DeleteFriendshipParams{
		ChatID:   chatID,
		FriendID: friendID,
	})
}

func (f *Friend) GetAprovedFriendships(ctx context.Context, chatID int64) ([]db.GetAprovedFriendshipsRow, error) {
	return f.db.GetAprovedFriendships(ctx, chatID)
}

func (f *Friend) GetPendingFriendships(ctx context.Context, chatID int64) ([]db.GetPendingFriendshipsRow, error) {
	return f.db.GetPendingFriendships(ctx, chatID)
}
func (f *Friend) UpdateFriendshipStatus(ctx context.Context, req db.UpdateFriendshipStatusParams) (db.Friend, error) {
	return f.db.UpdateFriendshipStatus(ctx, req)
}
