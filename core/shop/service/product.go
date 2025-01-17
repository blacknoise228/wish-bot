package service

import (
	"context"
	"log"
	"strconv"
	db "wish-bot/core/shop/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type Product struct {
	db  *db.Queries
	bot *tgbotapi.BotAPI
}

func NewProduct(db *db.Queries, bot *tgbotapi.BotAPI) Producter {
	return Product{
		db:  db,
		bot: bot,
	}
}

func (p Product) CreateProduct(chatID int64, data map[string]string) {

	price, err := strconv.ParseFloat(data["price"], 64)
	if err != nil {
		sendMessage(p.bot, chatID, "Некорректная цена")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	shop, err := p.db.GetShopAdminsByAdminID(context.Background(), chatID)
	if err != nil {
		sendMessage(p.bot, chatID, "Вы не являетесь администратором магазина!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	category, err := p.db.GetProductCategoryByName(context.Background(), data["category"])
	if err != nil {
		sendMessage(p.bot, chatID, "Такой категории не существует!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	p.db.CreateProduct(context.Background(), db.CreateProductParams{
		AdminID:     chatID,
		Name:        data["name"],
		Price:       price,
		Image:       data["image"],
		Description: data["desc"],
		ShopID:      shop.ShopID,
		CategoryID:  category.ID,
		Status:      1,
	})
	sendMessage(p.bot, chatID, "Товар успешно добавлен!")
}

func (p Product) UpdateProduct(chatID int64, data map[string]string) {

	price, err := strconv.ParseFloat(data["price"], 64)
	if err != nil {
		sendMessage(p.bot, chatID, "Некорректная цена")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	category, err := p.db.GetProductCategoryByName(context.Background(), data["category"])
	if err != nil {
		sendMessage(p.bot, chatID, "Такой категории не существует!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	id, err := uuid.Parse(data["id"])
	if err != nil {
		sendMessage(p.bot, chatID, "Некорректный ID товара!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	p.db.UpdateProduct(context.Background(), db.UpdateProductParams{
		AdminID:     chatID,
		Name:        data["name"],
		Price:       price,
		Image:       data["image"],
		Status:      1,
		CategoryID:  category.ID,
		ID:          id,
		Description: data["desc"],
	})

	sendMessage(p.bot, chatID, "Товар успешно обновлен!")
}

func (p Product) UpdateProductStatus(chatID int64, productID string, status string) {
	stat, err := p.db.GetProductStatusByName(context.Background(), status)
	if err != nil {
		sendMessage(p.bot, chatID, "Такого статуса не существует!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	shop, err := p.db.GetShopAdminsByAdminID(context.Background(), chatID)
	if err != nil {
		sendMessage(p.bot, chatID, "Вы не являетесь администратором магазина!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	id, err := uuid.Parse(productID)
	if err != nil {
		sendMessage(p.bot, chatID, "Некорректный ID товара!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}

	p.db.UpdateProductStatus(context.Background(), db.UpdateProductStatusParams{
		ID:     id,
		Status: stat.ID,
		ShopID: shop.ShopID,
	})

	sendMessage(p.bot, chatID, "Статус товара успешно обновлен!")
}
