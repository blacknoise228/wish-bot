package telegram

import (
	"log"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) sendMenuButton(chatID int64) {

	buttons := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Меню"),
	)
	msg := tgbotapi.NewMessage(chatID, "Меню")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}

func (t *Telegram) sendInlineMenu(chatID int64) {
	t.sendMenuButton(chatID)
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Регистрация администратора", "register"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить aдминистратора", "delete_admin"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать магазин", "edit_shop"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Заказы магазина", " shop_orders"),
			tgbotapi.NewInlineKeyboardButtonData("Заказы администратора", "admin_orders"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить товар", "add_product"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Товары", "products"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = buttons

	m, err := t.Bot.Send(msg)
	if err != nil {
		log.Println(errornator.CustomError(err.Error()))
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
	delete(lastMessageID, chatID)
	lastMessageID[chatID] = m.MessageID
}
