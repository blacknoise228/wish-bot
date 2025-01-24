package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackProductsHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	switch query.Data {
	case "product_flowers":
		go t.deleteLastMessage(chatID)
		t.Service.GetProductsByCategories(chatID, 1)

	case "product_clothes":
		go t.deleteLastMessage(chatID)
		t.Service.GetProductsByCategories(chatID, 2)

	case "product_electronics":
		go t.deleteLastMessage(chatID)
		t.Service.GetProductsByCategories(chatID, 3)

	case "product_toys":
		go t.deleteLastMessage(chatID)
		t.Service.GetProductsByCategories(chatID, 4)

	case "product_accessories":
		go t.deleteLastMessage(chatID)
		t.Service.GetProductsByCategories(chatID, 5)

	case "catalog":
		go t.deleteLastMessage(chatID)
		t.productsButton(chatID)

	}
}

func (t *Telegram) productsButton(chatID int64) {
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Цветы", "product_flowers"),
			tgbotapi.NewInlineKeyboardButtonData("Одежда", "product_clothes"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Электроника", "product_electronics"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Игрушки", "product_toys"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Аксессуары", "product_accessories"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите категорию: ")
	msg.ReplyMarkup = buttons

	m, err := t.Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
	delete(LastMessageID, chatID)
	LastMessageID[chatID] = m.MessageID

}
