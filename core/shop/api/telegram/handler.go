package telegram

import (
	"log"
	"wish-bot/core/shop/api/telegram/state"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	lastMessageID = make(map[int64]int)
	messageData   = make(map[int64]map[string]string)
)

func (t *Telegram) handleCallback(query *tgbotapi.CallbackQuery) {

	if _, exists := messageData[query.Message.Chat.ID]; !exists {
		messageData[query.Message.Chat.ID] = make(map[string]string)
	}

	if _, exists := lastMessageID[query.Message.Chat.ID]; !exists {
		lastMessageID = make(map[int64]int)
	}

	t.callbackProductHandler(query)
	t.callbackShopHandler(query)
	t.callbackOrderHandler(query)

	log.Println(query.Message.Chat.ID, query.Data)

	callback := tgbotapi.NewCallback(query.ID, "Ждем...")
	if _, err := t.Bot.Request(callback); err != nil {
		log.Println("Ошибка при обработке CallbackQuery:", err)
		log.Println(errornator.CustomError(err.Error()))
	}
}

func (t *Telegram) handleMessage(message *tgbotapi.Message) {

	userstate := state.GetUserState(message.Chat.ID)

	if _, exists := messageData[message.Chat.ID]; !exists {
		messageData[message.Chat.ID] = make(map[string]string)
	}

	if _, exists := lastMessageID[message.Chat.ID]; !exists {
		lastMessageID = make(map[int64]int)
	}

	t.createProductHandler(message, userstate)
	t.updateProductHandler(message, userstate)
	t.mesageShopHandler(message, userstate)
	t.messageOrderHandler(message, userstate)

	switch message.Text {
	case "/start":
		t.sendMenuButton(message.Chat.ID)
	case "Меню":
		state.ClearUserState(message.Chat.ID)
		t.sendInlineMenu(message.Chat.ID)
	}
}
