package telegram

import (
	"log"

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
			tgbotapi.NewInlineKeyboardButtonData("Регистрация", "register"),
			tgbotapi.NewInlineKeyboardButtonData("Удалить пользователя", "delete_user"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Редактировать никнейм", "edit_nickname"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Друзья", "friends"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить желания", "add_wish"),
			tgbotapi.NewInlineKeyboardButtonData("Мои желания", "my_wishes"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите действие: ")
	msg.ReplyMarkup = buttons

	m, err := t.Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
	delete(lastMessageID, chatID)
	lastMessageID[chatID] = m.MessageID
}

func (t *Telegram) sendSkipButton(chatID int64) {

	buttons := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Пропустить"),
	)
	msg := tgbotapi.NewMessage(chatID, "...")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}
