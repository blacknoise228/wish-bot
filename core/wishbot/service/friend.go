package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	db "wish-bot/core/wishbot/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Service) CreateFriendship(ctx context.Context, chatID int64, friendName string) error {
	user, err := t.DB.GetUserByUsername(ctx, friendName)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return err
	}
	_, err = t.DB.CreateFriendship(ctx, db.CreateFriendshipParams{
		ChatID:   chatID,
		FriendID: user.ChatID,
	})
	if err != nil {
		log.Println(err)
		t.sendMessage(chatID, "Ошибка отправки запроса дружбы")
		return err
	}
	user2, _ := t.DB.GetUser(ctx, chatID)

	id := strconv.FormatInt(chatID, 10)

	button := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Подтвердить", "approve:"+id),
			tgbotapi.NewInlineKeyboardButtonData("Отклонить", "decline:"+id),
		),
	)
	msg := tgbotapi.NewMessage(user.ChatID, "Пользователь "+user2.Username+" отправил вам запрос дружбы")
	msg.ReplyMarkup = button
	m, err := t.Bot.Send(msg)
	if err != nil {
		return err
	}

	if _, ok := lastMessage[user.ChatID]; !ok {
		lastMessage[user.ChatID] = m.MessageID
	}

	t.sendMessage(chatID, "Запрос отправлен!")

	return nil
}

func (t *Service) GetUserFriends(ctx context.Context, chatID int64) {
	friends, err := t.DB.GetAprovedFriendships(ctx, chatID)
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
				tgbotapi.NewInlineKeyboardButtonData("Удалить друга", "delete_friend:"+strconv.Itoa(int(friend.FriendID))),
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

func (t *Service) GetPendingFriends(ctx context.Context, chatID int64) {
	friends, err := t.DB.GetPendingFriendships(ctx, chatID)
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

func (t *Service) DeleteFriend(ctx context.Context, chatID int64, friendID int64) {

	err := t.DB.DeleteFriendship(ctx, db.DeleteFriendshipParams{ChatID: chatID, FriendID: friendID})
	if err != nil {
		t.sendMessage(chatID, "Ошибка при удалении дружбы")
		return
	}

	err = t.DB.DeleteFriendship(ctx, db.DeleteFriendshipParams{ChatID: friendID, FriendID: chatID})
	if err != nil {
		t.sendMessage(chatID, "Ошибка при удалении дружбы")
		return
	}

	t.sendMessage(chatID, "Дружба успешно удалена")
}

func (t *Service) UpdateFriendshipStatus(senderID, chatID int64, status int32) {
	t.DB.UpdateFriendshipStatus(context.Background(), db.UpdateFriendshipStatusParams{
		ChatID:   senderID,
		FriendID: chatID,
		Status:   status,
	})

	t.deleteLastMessage(chatID)
}
