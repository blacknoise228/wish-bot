package telegram

import (
	"context"
	"log"
	"wish-bot/core/wishbot/api/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	LastMessageID = make(map[int64]int)
	userWishData  = make(map[int64]map[string]string)
)

func (t *Telegram) handleCallback(query *tgbotapi.CallbackQuery) {

	if _, exists := userWishData[query.Message.Chat.ID]; !exists {
		userWishData[query.Message.Chat.ID] = make(map[string]string)
	}

	log.Println(query.Message.Chat.ID, query.Data)

	t.callbackWishHandler(query)

	t.callbackFriendHandler(query)

	t.callbackUsersHandler(query)

	callback := tgbotapi.NewCallback(query.ID, "Ждем...")
	if _, err := t.Bot.Request(callback); err != nil {
		log.Println("Ошибка при обработке CallbackQuery:", err)
	}
}

func (t *Telegram) handleMessage(ctx context.Context, message *tgbotapi.Message) {

	states := state.GetUserState(message.Chat.ID)

	if _, exists := userWishData[message.Chat.ID]; !exists {
		userWishData[message.Chat.ID] = make(map[string]string)
	}

	t.messageFriendHandler(ctx, states, message)

	t.messageUsersHandler(ctx, states, message)

	t.messageWishHandler(states, message)

	switch message.Text {
	case "/start":
		t.tgService.StartMessageHandler(message)
		t.sendMenuButton(message.Chat.ID)
	case "Меню":
		state.ClearUserState(message.Chat.ID)
		t.sendInlineMenu(message.Chat.ID)
	}
}
