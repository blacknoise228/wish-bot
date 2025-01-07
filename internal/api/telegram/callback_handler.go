package telegram

import (
	"context"
	"log"
	"strconv"
	"strings"
	"wish-bot/internal/api/telegram/state"
	db "wish-bot/internal/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) handleCallback(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID
	data := query.Data
	log.Println(chatID, data)
	if strings.HasPrefix(data, "approve:") {
		strID := strings.TrimPrefix(data, "approve:")
		senderID, _ := strconv.ParseInt(strID, 10, 64)

		log.Println(senderID)
		t.services.Friend.UpdateFriendshipStatus(context.Background(), db.UpdateFriendshipStatusParams{
			ChatID:   senderID,
			FriendID: chatID,
			Status:   1,
		})

		t.sendMessage(chatID, "Запрос дружбы принят!")

		t.sendMessage(senderID, "Запрос дружбы принят!")

	} else if strings.HasPrefix(data, "decline:") {
		strID := strings.TrimPrefix(data, "decline:")
		senderID, _ := strconv.ParseInt(strID, 10, 64)

		t.services.Friend.UpdateFriendshipStatus(context.Background(), db.UpdateFriendshipStatusParams{
			ChatID:   senderID,
			FriendID: chatID,
			Status:   3,
		})

		t.sendMessage(chatID, "Запрос дружбы отклонен!")

		t.sendMessage(senderID, "Запрос дружбы отклонен!")
	}

	switch query.Data {
	case "menu":
		t.sendInlineMenu(chatID)
	case "register":
		state.SetUserState(chatID, state.CreateUserWaiting)
		t.sendMessage(chatID, createNicknameMessage)
	case "edit_nickname":
		state.SetUserState(chatID, state.UpdateUserWaiting)
		t.sendMessage(chatID, updateUserMessage)
	case "delete_user":
		t.sendMessage(chatID, deleteUserMessage)
		t.deleteButton(chatID)
	case "yes_delete":
		if err := t.deleteUserHandler(query); err != nil {
			log.Println("Deleting user error: ", err)
			t.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		t.sendMessage(chatID, "Пользователь успешно удален.")

	case "add_wish":
		t.sendMessage(chatID, AddWishMessage)
		t.sendMessage(chatID, "Введите описание желания")
		state.SetUserState(chatID, state.AddWishDesc)
	case "my_wishes":
		t.sendMessage(chatID, "Ваши желания: ")
		t.getMyWishes(chatID)
	case "friends":
		t.friendsButton(chatID)
	case "add_friend":
		t.sendMessage(chatID, "Введите имя пользователя:")
		state.SetUserState(chatID, state.AddFriendWait)
	case "user_wishes":
		t.sendMessage(chatID, "Введите имя пользователя:")
		state.SetUserState(chatID, state.GetUserWish)
	case "delete_friend":
		t.sendMessage(chatID, "Введите имя пользователя:")
		state.SetUserState(chatID, state.RemoveFriendWait)
	case "my_friends":
		t.sendMessage(chatID, "Ваши друзья:")
		t.getUserFriends(context.Background(), chatID)
	case "pending_friends":
		t.sendMessage(chatID, "Ваши запросы в друзья:")
		t.getPendingFriends(context.Background(), chatID)
	}

	callback := tgbotapi.NewCallback(query.ID, "Ждем...")
	if _, err := t.Bot.Request(callback); err != nil {
		log.Println("Ошибка при обработке CallbackQuery:", err)
	}
}

func (t *Telegram) deleteButton(chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить профиль", "yes_delete"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Вы уверены???")
	msg.ReplyMarkup = buttons

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}

func (t *Telegram) friendsButton(chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить друга", "add_friend"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить друга", "delete_friend"),
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

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}

func (t *Telegram) sendInlineMenu(chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Регистрация", "register"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить пользователя", "delete_user"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать никнейм", "edit_nickname"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Друзья", "friends"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить желания", "add_wish"),
			tgbotapi.NewInlineKeyboardButtonData("Мои желания", "my_wishes"),
			tgbotapi.NewInlineKeyboardButtonData("Редактировать желание", "edit_wish"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = buttons

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}
