package tgservice

import (
	"log"
	"wish-bot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TGService struct {
	Bot      *tgbotapi.BotAPI
	Services *service.Services
}

func NewTGService(tgBot *tgbotapi.BotAPI, services *service.Services) *TGService {
	return &TGService{
		Bot:      tgBot,
		Services: services,
	}
}

func (t *TGService) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
		return err
	}

	return nil
}
