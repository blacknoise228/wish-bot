package telegram

import (
	"context"
	"log"
	"wish-bot/core/wishbot/api/telegram/state"
	"wish-bot/core/wishbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackUsersHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	switch query.Data {

	case "register":
		go t.deleteLastMessage(chatID)
		state.SetUserState(chatID, state.CreateUserWaiting)
		t.sendMessage(chatID, service.CreateNicknameMessage)
	case "edit_nickname":
		go t.deleteLastMessage(chatID)
		state.SetUserState(chatID, state.UpdateUserWaiting)
		t.sendMessage(chatID, service.UpdateUserMessage)
	case "delete_user":
		t.sendMessage(chatID, service.DeleteUserMessage)
		t.deleteButton(chatID)
	case "yes_delete":
		go t.deleteLastMessage(chatID)
		if err := t.Service.DeleteUserHandler(query); err != nil {
			log.Println("Deleting user error: ", err)
			t.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		t.sendMessage(chatID, "Пользователь успешно удален.")

	}
}

func (t *Telegram) messageUsersHandler(ctx context.Context, states string, message *tgbotapi.Message) {

	switch states {

	case state.CreateUserWaiting:
		if message.Text != "Меню" {
			if err := t.Service.CreateUserHandler(ctx, message); err != nil {
				log.Println(err)
				return
			}
			t.sendInlineMenu(message.Chat.ID)
		}
		state.ClearUserState(message.Chat.ID)

	case state.UpdateUserWaiting:
		if err := t.Service.UpdateUserHandler(ctx, message); err != nil {
			log.Println(err)
			return
		}
		state.ClearUserState(message.Chat.ID)
		t.sendInlineMenu(message.Chat.ID)
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

	m, err := t.Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
	delete(LastMessageID, chatID)
	LastMessageID[chatID] = m.MessageID
}
