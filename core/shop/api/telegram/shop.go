package telegram

import (
	"log"
	"wish-bot/core/shop/api/telegram/state"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackShopHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	switch query.Data {

	case "register":
		go t.deleteLastMessage(chatID)
		state.SetUserState(chatID, state.AddAdmin)
		t.sendMessage(chatID, "Введите токен магазина:")

	case "delete_admin":
		go t.deleteLastMessage(chatID)
		t.deleteButton(chatID)

	case "yes_delete":
		go t.deleteLastMessage(chatID)
		t.Service.Shop.DeleteShopAdmin(chatID)
		t.sendInlineMenu(chatID)
	}
}

func (t *Telegram) mesageShopHandler(message *tgbotapi.Message, userstate string) {

	chatID := message.Chat.ID

	switch userstate {

	case state.AddAdmin:
		t.Service.Shop.RegisterShopAdmin(chatID, message.Text)
		state.ClearUserState(chatID)
	}
}

func (t *Telegram) deleteButton(chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить администратора", "yes_delete"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Вы уверены???")
	msg.ReplyMarkup = buttons

	m, err := t.Bot.Send(msg)
	if err != nil {
		log.Println(errornator.CustomError(err.Error()))
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
	delete(lastMessageID, chatID)
	lastMessageID[chatID] = m.MessageID
}
