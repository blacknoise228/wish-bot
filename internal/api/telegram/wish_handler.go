package telegram

import (
	"strings"
	"wish-bot/internal/api/telegram/state"
	tgservice "wish-bot/internal/tg-service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackWishHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	wishMap := userWishData[chatID]

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
		t.sendMessage(chatID, tgservice.AddWishMessage)
		t.sendMessage(chatID, "Введите описание желания")
		state.SetUserState(chatID, state.AddWishDesc)
	case "my_wishes":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Ваши желания: ")
		t.tgService.GetMyWishes(chatID)
	case "edit_wish":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Введите описание желания")
		state.SetUserState(chatID, state.UpdateWishDesc)
	}
	if strings.HasPrefix(query.Data, "update_wish:") {
		id := strings.TrimPrefix(query.Data, "update_wish:")
		wishMap["wish_id"] = id

		t.sendMessage(chatID, "Введите описание желания")

		state.SetUserState(chatID, state.UpdateWishDesc)
	}
	if strings.HasPrefix(query.Data, "remove_wish:") {
		id := strings.TrimPrefix(query.Data, "remove_wish:")
		t.tgService.DeleteWish(chatID, id)
	}
}

func (t *Telegram) messageWishHandler(states string, message *tgbotapi.Message) {

	chatID := message.Chat.ID

	wishMap := userWishData[chatID]

	switch states {

	case state.AddWishDesc:
		t.sendSkipButton(chatID)
		if message.Text != "Пропустить" {
			wishMap["desc"] = message.Text
		}
		state.SetUserState(chatID, state.AddWishLink)
		t.sendMessage(chatID, "Добавьте ссылку")
	case state.AddWishLink:
		t.sendSkipButton(chatID)
		if message.Text != "Пропустить" {
			wishMap["link"] = message.Text
		}
		state.SetUserState(chatID, state.AddWishStat)
		t.sendMessage(chatID, "Выберите тип желания:\n1. Публичный\n2. Приватный\nВведите число соответствующее типу желания.")
	case state.AddWishStat:
		wishMap["status"] = message.Text
		t.tgService.CreateWish(chatID, wishMap)
		state.ClearUserState(chatID)
		t.sendInlineMenu(chatID)
		t.sendMenuButton(chatID)
		clearWishMap(wishMap)

	case state.UpdateWishDesc:
		t.sendSkipButton(chatID)
		if message.Text != "Пропустить" {
			wishMap["desc"] = message.Text
		}
		state.SetUserState(chatID, state.UpdateWishLink)
		t.sendMessage(chatID, "Добавьте ссылку")
	case state.UpdateWishLink:
		t.sendSkipButton(chatID)
		if message.Text != "Пропустить" {
			wishMap["link"] = message.Text
		}
		state.SetUserState(chatID, state.UpdateWishStat)
		t.sendMessage(chatID, "Выберите тип желания:\n1. Публичный\n2. Приватный\nВведите число соответствующее типу желания.")
	case state.UpdateWishStat:
		wishMap["status"] = message.Text
		t.tgService.UpdateWish(chatID, wishMap)
		t.sendInlineMenu(chatID)
		t.sendMenuButton(chatID)
		clearWishMap(wishMap)
		state.ClearUserState(chatID)

	case state.GetUserWish:
		t.tgService.GetUserWishes(chatID, message.Text)
		state.ClearUserState(chatID)

	case state.DeleteWish:
		t.tgService.DeleteWish(chatID, message.Text)
		state.ClearUserState(chatID)
		t.sendInlineMenu(chatID)
	}
}

func clearWishMap(wishMap map[string]string) {
	for key := range wishMap {
		delete(wishMap, key)
	}
}
