package service

import (
	"context"
	"wish-bot/internal/config"
	db "wish-bot/internal/db/sqlc"
)

type Wish struct {
	db     *db.Queries
	config *config.Config
}

func NewWish(db *db.Queries, config *config.Config) Wisher {
	return &Wish{
		db:     db,
		config: config,
	}
}

func (w *Wish) CreateWish(ctx context.Context,
	req db.CreateWishParams) (db.Wish, error) {
	return w.db.CreateWish(ctx, req)
}

func (w *Wish) GetWishesForUser(ctx context.Context,
	chatID int64) ([]db.GetWishesForUserRow, error) {
	return w.db.GetWishesForUser(ctx, chatID)
}

func (w *Wish) GetUserWishes(ctx context.Context, chatID int64, friendID int64) []db.GetWishesPublicRow {
	wishes, err := w.db.GetWishesPublic(ctx, db.GetWishesPublicParams{
		ChatID:   friendID,
		ChatID_2: chatID,
	})
	if err != nil {
		return nil
	}
	return wishes
}

func (w *Wish) DeleteWish(ctx context.Context, chatID int64, wishID int) error {
	_, err := w.db.GetWish(ctx, db.GetWishParams{
		ChatID: chatID,
		ID:     int32(wishID),
	})
	if err != nil {
		return err
	}
	return w.db.DeleteWish(ctx, db.DeleteWishParams{
		ID:     int32(wishID),
		ChatID: chatID,
	})
}

func (w *Wish) UpdateWish(ctx context.Context, req db.UpdateWishParams) (db.Wish, error) {
	return w.db.UpdateWish(ctx, req)
}

func (w *Wish) GetWish(ctx context.Context, chatID int64, wishID int) (db.Wish, error) {
	return w.db.GetWish(ctx, db.GetWishParams{
		ChatID: chatID,
		ID:     int32(wishID),
	})
}
