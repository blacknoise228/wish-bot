package service

import (
	"log"
	db "wish-bot/core/wishbot/db/sqlc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var lastMessage = make(map[int64]int)

type Service struct {
	Bot *tgbotapi.BotAPI
	DB  *db.Queries
}

func NewService(tgBot *tgbotapi.BotAPI, db *db.Queries) *Service {
	return &Service{
		Bot: tgBot,
		DB:  db,
	}
}

func (t *Service) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
		return err
	}

	return nil
}

func (t *Service) deleteLastMessage(chatID int64) {
	msg := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: lastMessage[chatID],
	}

	if _, err := t.Bot.Request(msg); err != nil {
		log.Printf("Ошибка удаления сообщения: %v\n", err)
	}
	delete(lastMessage, chatID)
}
