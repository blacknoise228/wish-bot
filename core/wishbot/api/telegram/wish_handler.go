package telegram

import (
	"strings"
	"wish-bot/core/wishbot/api/telegram/state"
	"wish-bot/core/wishbot/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackWishHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	switch query.Data {
	case "user_wishes":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Введите имя пользователя:")
		state.SetUserState(chatID, state.GetUserWish)
	case "delete_wish":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Введите ID желания:")
		state.SetUserState(chatID, state.DeleteWish)
	case "add_wish":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, service.AddWishMessage)
		t.sendMessage(chatID, "Введите описание желания")
		t.sendSkipButton(chatID)
		state.SetUserState(chatID, state.AddWishDesc)
	case "my_wishes":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Ваши желания: ")
		t.Service.GetMyWishes(chatID)
	case "edit_wish":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Введите описание желания")
		state.SetUserState(chatID, state.UpdateWishDesc)
	}
	if strings.HasPrefix(query.Data, "update_wish:") {
		id := strings.TrimPrefix(query.Data, "update_wish:")
		t.Service.UpdateWish(chatID, id)
	}
	if strings.HasPrefix(query.Data, "remove_wish:") {
		id := strings.TrimPrefix(query.Data, "remove_wish:")
		t.Service.DeleteWish(chatID, id)
	}
}

func (t *Telegram) messageWishHandler(states string, message *tgbotapi.Message) {

	chatID := message.Chat.ID

	switch states {

	case state.GetUserWish:
		t.Service.GetUserWishes(chatID, message.Text)
		state.ClearUserState(chatID)

	case state.DeleteWish:
		t.Service.DeleteWish(chatID, message.Text)
		state.ClearUserState(chatID)
		t.sendInlineMenu(chatID)
	}
}
