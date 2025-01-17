package service

import (
	"context"
	"log"
	db "wish-bot/core/shop/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Shop struct {
	bot *tgbotapi.BotAPI
	db  *db.Queries
}

func NewShop(bot *tgbotapi.BotAPI, db *db.Queries) Shoper {
	return &Shop{
		bot: bot,
		db:  db,
	}
}

func (s *Shop) RegisterShopAdmin(chatID int64, token string) {

	shop, err := s.db.GetShopByToken(context.Background(), token)
	if err != nil {
		log.Println(errornator.CustomError(err.Error()))
		sendMessage(s.bot, chatID, "Неверный токен магазина! Такого магазина не существует!")
		return
	}

	if err := s.db.CreateShopAdmin(context.Background(), db.CreateShopAdminParams{
		AdminID: chatID,
		ShopID:  shop.ID,
	}); err != nil {
		log.Println(errornator.CustomError(err.Error()))
		sendMessage(s.bot, chatID, "Ошибка при добавлении администратора магазина!\nВозможно, вы уже являетесь администратором магазина!")
		return
	}

	sendMessage(s.bot, chatID, "Вы успешно зарегистрировались как администратор магазина!")
}

func (s *Shop) DeleteShopAdmin(chatID int64) {

	if err := s.db.DeleteShopAdmin(context.Background(), chatID); err != nil {
		sendMessage(s.bot, chatID, "Ошибка при удалении администратора магазина!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	sendMessage(s.bot, chatID, "Вы успешно удалены из администраторов магазина!")
}
