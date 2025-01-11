package service

import (
	"log"
	db "wish-bot/core/shop/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Service struct {
	bot *tgbotapi.BotAPI
	db  *db.Queries
}

func NewService(tgBot *tgbotapi.BotAPI, db *db.Queries) *Service {
	return &Service{
		bot: tgBot,
		db:  db,
	}
}

func (t *Service) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := t.bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
		return err
	}

	return nil
}
