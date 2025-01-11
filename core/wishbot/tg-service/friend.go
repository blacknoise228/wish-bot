package tgservice

import (
	"context"
	"fmt"
	"log"
	"strconv"
	db "wish-bot/core/wishbot/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *TGService) CreateFriendship(ctx context.Context, chatID int64, friendName string) error {
	user, err := t.Services.User.GetUserByUsername(ctx, friendName)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return err
	}
	_, err = t.Services.Friend.CreateFriendship(ctx, db.CreateFriendshipParams{
		ChatID:   chatID,
		FriendID: user.ChatID,
	})
	if err != nil {
		log.Println(err)
		t.sendMessage(chatID, "Ошибка отправки запроса дружбы")
		return err
	}
	user2, _ := t.Services.User.GetUser(ctx, chatID)

	id := strconv.FormatInt(chatID, 10)

	button := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить", "approve:"+id),
			tgbotapi.NewInlineKeyboardButtonData("Отклонить", "decline:"+id),
		),
	)
	msg := tgbotapi.NewMessage(user.ChatID, "Пользователь "+user2.Username+" отправил вам запрос дружбы")
	msg.ReplyMarkup = button
	_, err = t.Bot.Send(msg)
	if err != nil {
		return err
	}

	t.sendMessage(user.ChatID, "Пользователь отправил вам запрос дружбы")

	t.sendMessage(chatID, "Запрос отправлен!")

	return nil
}

func (t *TGService) GetUserFriends(ctx context.Context, chatID int64) {
	friends, err := t.Services.Friend.GetAprovedFriendships(ctx, chatID)
	if err != nil {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	if len(friends) == 0 {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	for _, friend := range friends {

		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Удалить друга", "delete_friend:+"+strconv.Itoa(int(friend.FriendID))),
				tgbotapi.NewInlineKeyboardButtonData("Желание друга", "get_wishes:"+friend.Username),
			),
		)

		resp := fmt.Sprintf("Имя друга: %v", friend.Username)

		msg := tgbotapi.NewMessage(chatID, resp)
		msg.ReplyMarkup = buttons
		_, err := t.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", err)
		}
	}
}

func (t *TGService) GetPendingFriends(ctx context.Context, chatID int64) {
	friends, err := t.Services.Friend.GetPendingFriendships(ctx, chatID)
	if err != nil {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	if len(friends) == 0 {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	for _, friend := range friends {

		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Удалить друга", "delete_friend:+"+strconv.Itoa(int(friend.FriendID))),
				tgbotapi.NewInlineKeyboardButtonData("Желание друга", "get_wish:"+friend.Username),
			),
		)

		resp := fmt.Sprintf("Имя друга: %v\nСтатус: %v", friend.Username, friend.StatusName)

		msg := tgbotapi.NewMessage(chatID, resp)
		msg.ReplyMarkup = buttons
		_, err := t.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", err)
		}
	}
}

func (t *TGService) DeleteFriend(ctx context.Context, chatID int64, friendID int64) {

	err := t.Services.Friend.DeleteFriendship(ctx, chatID, friendID)
	if err != nil {
		t.sendMessage(chatID, "Ошибка при удалении дружбы")
		return
	}
	t.sendMessage(chatID, "Дружба успешно удалена")
}
