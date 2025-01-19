package telegram

import (
	"context"
	"log"
	"strconv"
	"strings"
	"wish-bot/core/wishbot/api/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackFriendHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	switch query.Data {
	case "friends":
		go t.deleteLastMessage(chatID)
		t.friendsButton(chatID)
	case "add_friend":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Введите имя пользователя:")
		state.SetUserState(chatID, state.AddFriendWait)
	case "my_friends":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Ваши друзья:")
		t.Service.GetUserFriends(context.Background(), chatID)
	case "pending_friends":
		t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Ваши запросы в друзья:")
		t.Service.GetPendingFriends(context.Background(), chatID)
	}

	t.approveFriendHandler(query.Data, chatID)

	if strings.HasPrefix(query.Data, "delete_friend:") {
		strID := strings.TrimPrefix(query.Data, "delete_friend:")
		senderID, _ := strconv.ParseInt(strID, 10, 64)

		t.Service.DeleteFriend(context.Background(), chatID, senderID)
	}
}

func (t *Telegram) messageFriendHandler(ctx context.Context, states string, message *tgbotapi.Message) {

	switch states {

	case state.AddFriendWait:
		t.Service.CreateFriendship(ctx, message.Chat.ID, message.Text)
		state.ClearUserState(message.Chat.ID)
	}
}

func (t *Telegram) friendsButton(chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить друга", "add_friend"),
			tgbotapi.NewInlineKeyboardButtonData("Мои заявки", "pending_friends"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Мои друзья", "my_friends"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Желания пользователя", "user_wishes"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = buttons

	m, err := t.Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
	delete(LastMessageID, chatID)
	LastMessageID[chatID] = m.MessageID

}

func (t *Telegram) approveFriendHandler(data string, chatID int64) {
	if strings.HasPrefix(data, "approve:") {
		strID := strings.TrimPrefix(data, "approve:")
		senderID, _ := strconv.ParseInt(strID, 10, 64)

		log.Println(senderID)

		t.Service.UpdateFriendshipStatus(senderID, chatID, 1)

		t.sendMessage(chatID, "Запрос дружбы принят!")

		t.sendMessage(senderID, "Запрос дружбы принят!")

	} else if strings.HasPrefix(data, "decline:") {
		strID := strings.TrimPrefix(data, "decline:")
		senderID, _ := strconv.ParseInt(strID, 10, 64)

		t.Service.DeleteFriend(context.Background(), senderID, chatID)

		t.sendMessage(chatID, "Запрос дружбы отклонен!")

		t.sendMessage(senderID, "Запрос дружбы отклонен!")
	}
}
