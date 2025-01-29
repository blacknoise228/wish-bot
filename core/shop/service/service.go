package service

import (
	"log"
	db "wish-bot/core/shop/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type Service struct {
	bot     *tgbotapi.BotAPI
	db      *db.Queries
	Product Producter
	Shop    Shoper
	Order   Order
}

func NewService(tgBot *tgbotapi.BotAPI, db *db.Queries) *Service {
	return &Service{
		bot:     tgBot,
		db:      db,
		Product: NewProduct(db, tgBot),
		Shop:    NewShop(tgBot, db),
		Order:   NewOrderService(db, tgBot),
	}
}

type Producter interface {
	CreateProduct(chatID int64, data map[string]string)
	UpdateProduct(chatID int64, data map[string]string)
	UpdateProductStatus(chatID int64, productID string, status string)
}

type Shoper interface {
	RegisterShopAdmin(chatID int64, token string)
	DeleteShopAdmin(chatID int64)
}

type Order interface {
	GetAdminOrders(chatID int64)
	GetShopOrders(chatID int64)
	UpdateOrderStatus(chatID int64, orderID uuid.UUID, status int32)
	SendPayLink(chatID int64, orderID uuid.UUID, link string)
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
		log.Println(errornator.CustomError(err.Error()))
		return err
	}
	return nil
}
