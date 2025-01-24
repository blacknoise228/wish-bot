package telegram

import (
	"strings"
	"wish-bot/core/wishbot/api/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func (t *Telegram) callbackWishHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	switch query.Data {
	case "user_wishes":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Введите имя пользователя:")
		state.SetUserState(chatID, state.GetUserWish)

	case "my_wishes":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Ваши желания: ")
		t.Service.GetMyWishes(chatID)

	}
	if strings.HasPrefix(query.Data, "update_wish:") {
		id := strings.TrimPrefix(query.Data, "update_wish:")
		t.Service.UpdateWish(chatID, id)
	}
	if strings.HasPrefix(query.Data, "remove_wish:") {
		id := strings.TrimPrefix(query.Data, "remove_wish:")
		t.Service.DeleteWish(chatID, id)
	}

	if strings.HasPrefix(query.Data, "get_wishes:") {
		userName := strings.TrimPrefix(query.Data, "get_wishes:")

		t.Service.GetUserWishes(chatID, userName)
	}

	if strings.HasPrefix(query.Data, "add_wish_private:") {
		productID := strings.TrimPrefix(query.Data, "add_wish_private:")

		id := uuid.MustParse(productID)

		t.Service.CreateWish(chatID, id, 2)
	}

	if strings.HasPrefix(query.Data, "add_wish_public:") {
		productID := strings.TrimPrefix(query.Data, "add_wish_public:")

		id := uuid.MustParse(productID)

		t.Service.CreateWish(chatID, id, 1)
	}
}

func (t *Telegram) messageWishHandler(states string, message *tgbotapi.Message) {

	chatID := message.Chat.ID

	switch states {

	case state.GetUserWish:
		t.Service.GetUserWishes(chatID, message.Text)
		state.ClearUserState(chatID)
	}

}
