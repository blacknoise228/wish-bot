package telegram

import (
	"context"
	"fmt"
	"log"
	"strconv"
	db "wish-bot/internal/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) createFriendship(ctx context.Context, chatID int64, friendName string) error {
	user, err := t.services.User.GetUserByUsername(ctx, friendName)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return err
	}
	_, err = t.services.Friend.CreateFriendship(ctx, db.CreateFriendshipParams{
		ChatID:   chatID,
		FriendID: user.ChatID,
	})
	if err != nil {
		log.Println(err)
		t.sendMessage(chatID, "Ошибка отправки запроса дружбы")
		return err
	}
	user2, _ := t.services.User.GetUser(ctx, chatID)

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

	return nil
}

func (t *Telegram) getUserFriends(ctx context.Context, chatID int64) {
	friends, err := t.services.Friend.GetAprovedFriendships(ctx, chatID)
	if err != nil {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	if len(friends) == 0 {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	for _, friend := range friends {
		msg := fmt.Sprintf("Имя друга: %v\n", friend.Username)
		t.sendMessage(chatID, msg)
	}
}

func (t *Telegram) getPendingFriends(ctx context.Context, chatID int64) {
	friends, err := t.services.Friend.GetPendingFriendships(ctx, chatID)
	if err != nil {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	if len(friends) == 0 {
		t.sendMessage(chatID, "У вас нет друзей!\nAXAXAXAX")
		return
	}
	for _, friend := range friends {
		msg := fmt.Sprintf("Имя друга: %v\nСтатус: %v", friend.Username, friend.StatusName)
		t.sendMessage(chatID, msg)
	}
}

func (t *Telegram) deleteFriend(ctx context.Context, chatID int64, friendName string) {
	user, err := t.services.User.GetUserByUsername(ctx, friendName)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return
	}

	err = t.services.Friend.DeleteFriendship(ctx, chatID, user.ChatID)
	if err != nil {
		t.sendMessage(chatID, "Ошибка при удалении дружбы")
		return
	}
	t.sendMessage(chatID, "Дружба успешно удалена")
}
