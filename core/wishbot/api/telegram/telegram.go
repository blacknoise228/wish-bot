package telegram

import (
	"context"
	"log"
	"sync"
	"wish-bot/core/wishbot/config"
	"wish-bot/core/wishbot/service"
	tgservice "wish-bot/core/wishbot/tg-service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Bot       *tgbotapi.BotAPI
	tgService *tgservice.TGService
}

func NewTelegram(cfg *config.Config, services *service.Services) *Telegram {

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	tgservice := tgservice.NewTGService(bot, services)

	return &Telegram{
		Bot:       bot,
		tgService: tgservice,
	}
}

func (t *Telegram) StartBot(ctx context.Context) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.Bot.GetUpdatesChan(u)

	var wg sync.WaitGroup

	for update := range updates {
		wg.Add(1)
		go func(update tgbotapi.Update) {
			defer wg.Done()
			t.handleUpdate(ctx, update)
		}(update)
	}

	wg.Wait()

}

func (t *Telegram) handleUpdate(ctx context.Context, update tgbotapi.Update) {

	if update.CallbackQuery != nil {
		t.handleCallback(update.CallbackQuery)
	}
	if update.Message != nil {
		t.handleMessage(ctx, update.Message)
	}
}

func (t *Telegram) sendMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
		return err
	}

	return nil
}

func (t *Telegram) deleteLastMessage(chatID int64) {
	msg := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: LastMessageID[chatID],
	}

	if _, err := t.Bot.Request(msg); err != nil {
		log.Printf("Ошибка удаления сообщения: %v\n", err)
	}
	delete(LastMessageID, chatID)
}
