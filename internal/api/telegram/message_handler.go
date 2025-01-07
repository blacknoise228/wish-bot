package telegram

import (
	"context"
	"log"
	"wish-bot/internal/api/telegram/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userWishData = make(map[int64]map[string]string)

func (t *Telegram) handleMessage(ctx context.Context, message *tgbotapi.Message) {
	chatID := message.Chat.ID
	states := state.GetUserState(chatID)

	if _, exists := userWishData[chatID]; !exists {
		userWishData[chatID] = make(map[string]string)
	}
	wishMap := userWishData[chatID]

	switch message.Text {
	case "/start":
		t.startMessageHandler(message)
		t.sendMenuButton(chatID)
	}
	switch states {
	case state.CreateUserWaiting:
		if err := t.createUserHandler(ctx, message); err != nil {
			log.Println(err)
			return
		}
		state.ClearUserState(chatID)
		t.sendInlineMenu(chatID)
	case state.UpdateUserWaiting:
		if err := t.updateUserHandler(ctx, message); err != nil {
			log.Println(err)
			return
		}
		state.ClearUserState(chatID)
		t.sendInlineMenu(chatID)

	case state.AddWishDesc:
		wishMap["desc"] = message.Text
		state.SetUserState(chatID, state.AddWishLink)
		t.sendMessage(chatID, "Добавьте ссылку")
	case state.AddWishLink:
		wishMap["link"] = message.Text
		state.SetUserState(chatID, state.AddWishStat)
		t.sendMessage(chatID, "Выберите тип желания:\n1. Публичный\n2. Приватный\nВведите число соответствующее типу желания.")
	case state.AddWishStat:
		wishMap["status"] = message.Text
		t.createWish(chatID, wishMap)
		state.ClearUserState(chatID)
		t.sendInlineMenu(chatID)
		clearWishMap(wishMap)
	case state.GetUserWish:
		t.getUserWishes(chatID, message.Text)
		state.ClearUserState(chatID)
	case state.AddFriendWait:
		t.createFriendship(ctx, chatID, message.Text)
		state.ClearUserState(chatID)
	case state.RemoveFriendWait:
		t.deleteFriend(ctx, chatID, message.Text)
		state.ClearUserState(chatID)
	}
}

func clearWishMap(wishMap map[string]string) {
	for key := range wishMap {
		delete(wishMap, key)
	}
}

func (t *Telegram) sendMenuButton(chatID int64) {
	buttons := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Меню"),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}
