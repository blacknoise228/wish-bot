package telegram

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	db "wish-bot/internal/db/sqlc"

	"github.com/jackc/pgx/v5"
)

func (t *Telegram) createWish(chatID int64, wishMap map[string]string) {
	stat, err := strconv.Atoi(wishMap["status"])
	if err != nil {
		t.sendMessage(chatID, "Некорректное значение статуса")
		return
	}

	_, err = t.services.Wish.CreateWish(context.Background(), db.CreateWishParams{
		ChatID:      chatID,
		Description: wishMap["desc"],
		Link:        wishMap["link"],
		Status:      int32(stat),
	})
	if err != nil {
		t.sendMessage(chatID, "Ошибка при добавлении желания!")
	}
	t.sendMessage(chatID, "Желание добавлено!")
}

func (t *Telegram) getMyWishes(chatID int64) {
	wishes, err := t.services.Wish.GetWishesForUser(context.Background(), chatID)
	if err != nil {
		t.sendMessage(chatID, "Ошибка при получении желаний")
		return
	}
	for _, wish := range wishes {
		msg := fmt.Sprintf("ID:%d\nИмя пользователя: %v\nОписание: %v\nСсылка: %v\nДата создания: %v\nСтатус: %v",
			wish.ID, wish.Username, wish.Description, wish.Link, wish.CreatedAt, wish.StatusName)
		t.sendMessage(chatID, msg)
	}
}

func (t *Telegram) getUserWishes(chatID int64, friendUsername string) {
	friend, err := t.services.User.GetUserByUsername(context.Background(), friendUsername)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return
	}
	wishes := t.services.Wish.GetUserWishes(context.Background(), chatID, friend.ChatID)
	if len(wishes) == 0 {
		t.sendMessage(chatID, "У пользователя нет желаний!")
		return
	}
	for _, wish := range wishes {
		msg := fmt.Sprintf("Имя пользователя: %v\nОписание: %v\nСсылка: %v\nДата создания: %v\nСтатус: %v",
			wish.Username, wish.Description, wish.Link, wish.CreatedAt, wish.StatusName)
		t.sendMessage(chatID, msg)
	}
}

func (t *Telegram) deleteWish(chatID int64, wishID string) {
	wish, err := strconv.Atoi(wishID)
	if err != nil {
		t.sendMessage(chatID, "Некорректное значение ID")
		return
	}
	err = t.services.Wish.DeleteWish(context.Background(), chatID, wish)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			t.sendMessage(chatID, "Такого желания не существует!")
			return
		}
		t.sendMessage(chatID, "Ошибка при удалении желания!")
		return
	}
	t.sendMessage(chatID, "Желание удалено!")
}
