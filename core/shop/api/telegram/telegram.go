package telegram

import (
	"context"
	"log"
	"sync"
	"wish-bot/core/shop/config"
	db "wish-bot/core/shop/db/sqlc"
	"wish-bot/core/shop/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Bot     *tgbotapi.BotAPI
	Service *service.Service
}

func NewTelegram(cfg *config.Config, db *db.Queries) *Telegram {

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	service := service.NewService(bot, db)

	return &Telegram{
		Bot:     bot,
		Service: service,
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
		MessageID: lastMessageID[chatID],
	}

	if _, err := t.Bot.Request(msg); err != nil {
		log.Printf("Ошибка удаления сообщения: %v\n", err)
	}
	delete(lastMessageID, chatID)
}
