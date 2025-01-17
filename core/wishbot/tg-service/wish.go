package tgservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	db "wish-bot/core/wishbot/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func (t *TGService) CreateWish(chatID int64, wishMap map[string]string) {
	stat, err := strconv.Atoi(wishMap["status"])
	if err != nil {
		t.sendMessage(chatID, "Некорректное значение статуса")
		return
	}

	_, err = t.Services.Wish.CreateWish(context.Background(), db.CreateWishParams{
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

func (t *TGService) GetMyWishes(chatID int64) {
	wishes, err := t.Services.Wish.GetWishesForUser(context.Background(), chatID)
	if err != nil {
		t.sendMessage(chatID, "Ошибка при получении желаний")
		return
	}
	if len(wishes) == 0 {
		t.sendMessage(chatID, "У вас нет желаний!")
		return
	}

	for _, wish := range wishes {

		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Редактировать желание", "update_wish:"+strconv.Itoa(int(wish.ID))),
				tgbotapi.NewInlineKeyboardButtonData("Удалить желание", "remove_wish:+"+strconv.Itoa(int(wish.ID))),
			),
		)
		resp := fmt.Sprintf("ID:%d\nИмя пользователя: %v\nОписание: %v\nСсылка: %v\nДата создания: %v\nСтатус: %v",
			wish.ID, wish.Username, wish.Description, wish.Link, wish.CreatedAt, wish.StatusName)
		msg := tgbotapi.NewMessage(chatID, resp)
		msg.ReplyMarkup = buttons
		_, err := t.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", err)
		}
	}
}

func (t *TGService) GetUserWishes(chatID int64, friendUsername string) {
	friend, err := t.Services.User.GetUserByUsername(context.Background(), friendUsername)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return
	}
	wishes := t.Services.Wish.GetUserWishes(context.Background(), chatID, friend.ChatID)
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

func (t *TGService) DeleteWish(chatID int64, wishID string) {

	wish, err := strconv.Atoi(wishID)
	if err != nil {
		t.sendMessage(chatID, "Некорректное значение ID")
		return
	}

	err = t.Services.Wish.DeleteWish(context.Background(), chatID, wish)
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

func (t *TGService) UpdateWish(chatID int64, wishMap map[string]string) {

	stat, err := strconv.Atoi(wishMap["status"])
	if err != nil {
		t.sendMessage(chatID, "Некорректное значение статуса")
		return
	}

	id, _ := strconv.Atoi(wishMap["wish_id"])

	wish, err := t.Services.Wish.GetWish(context.Background(), chatID, id)
	if err != nil {
		t.sendMessage(chatID, "Такого желания не существует!")
		return
	}
	if wish.ChatID != chatID {
		t.sendMessage(chatID, "Ты не можете обновить чужое желание!")
		return
	}

	updWish := db.UpdateWishParams{
		ID:          int32(id),
		ChatID:      chatID,
		Description: wishMap["desc"],
		Link:        wishMap["link"],
		Status:      int32(stat),
	}

	if wishMap["desc"] == "" {
		updWish.Description = wish.Description
	}
	if wishMap["link"] == "" {
		updWish.Link = wish.Link
	}

	_, err = t.Services.Wish.UpdateWish(context.Background(), updWish)
	if err != nil {
		t.sendMessage(chatID, "Ошибка при обновлении желания!")
	}

	t.sendMessage(chatID, "Желание обновлено!")
}
