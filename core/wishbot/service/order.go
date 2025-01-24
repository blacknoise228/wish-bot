package service

import (
	"context"
	"fmt"
	"log"
	db "wish-bot/core/wishbot/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func (t *Service) CreateOrder(customerID int64, customerLogin string, wishID int32) {
	wish, err := t.DB.GetWishByID(context.Background(), wishID)
	if err != nil {
		log.Println(errornator.CustomError("Такого желания не существует!" + err.Error()))
		t.sendMessage(customerID, "Ошибка оформления заказа!")
		return
	}

	product, err := t.DB.GetProductByID(context.Background(), wish.ProductID)
	if err != nil {
		log.Println(errornator.CustomError("Такого товара не существует!" + err.Error()))
		t.sendMessage(customerID, "Ошибка оформления заказа!")
		return
	}

	admin, err := t.DB.GetRandomAdminByShopID(context.Background(), product.ShopID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении администратора!" + err.Error()))
		t.sendMessage(customerID, "Ошибка оформления заказа!")
		return
	}

	if _, err = t.DB.CreateOrder(context.Background(), db.CreateOrderParams{
		CustomerID:    customerID,
		CustomerLogin: customerLogin,
		ProductID:     product.ID,
		Price:         product.Price,
		Status:        1,
		AdminID:       admin.AdminID,
		ShopID:        product.ShopID,
		ConsigneeID:   wish.ChatID,
	}); err != nil {
		log.Println(errornator.CustomError("Ошибка оформления заказа!" + err.Error()))
		t.sendMessage(customerID, "Ошибка оформления заказа!")
	}
	t.sendMessage(customerID, "Заказ успешно оформлен!")
}

func (t *Service) UpdateOrderStatus(chatID int64, orderID uuid.UUID, status int32) {
	err := t.DB.UpdateOrderStatus(context.Background(), db.UpdateOrderStatusParams{
		ID:         orderID,
		CustomerID: chatID,
		Status:     status,
	})
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при обновлении статуса заказа!" + err.Error()))
		t.sendMessage(chatID, "Ошибка при обновлении статуса заказа!")
		return
	}
	t.sendMessage(chatID, "Статус заказа успешно обновлен!")
}

func (t *Service) GetOrderStatus(chatID int64, orderID uuid.UUID) {
	order, err := t.DB.GetOrder(context.Background(), orderID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении статуса заказа!" + err.Error()))
		t.sendMessage(chatID, "Ошибка при получении статуса заказа!")
		return
	}

	stat, err := t.DB.GetDimOrderStatusByID(context.Background(), order.Status)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении статуса заказа!" + err.Error()))
		t.sendMessage(chatID, "Ошибка при получении статуса заказа!")
		return
	}

	t.sendMessage(chatID, "Статус заказа: "+stat.StatusName)
}

func (t *Service) GetOrders(chatID int64) {
	orders, err := t.DB.GetOrdersByCustomer(context.Background(), chatID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении заказов!" + err.Error()))
		t.sendMessage(chatID, "Ошибка при получении заказов!")
		return
	}

	for _, order := range orders {

		stat, err := t.DB.GetDimOrderStatusByID(context.Background(), order.Status)
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении заказов!" + err.Error()))
			t.sendMessage(chatID, "Ошибка при получении заказов!")
			return
		}

		product, err := t.DB.GetProductByID(context.Background(), order.ProductID)
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении заказa!" + err.Error()))
			t.sendMessage(chatID, "Ошибка при получении заказa!")
			continue
		}

		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Отменить", "cancel_order:"+order.ID.String()),
			),
		)
		resp := fmt.Sprintf("ID заказа: %v\nОписание: %v\nЦена: %v\nСтатус: %v",
			order.ID, product.Description, product.Price, stat.StatusName)
		msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(product.Image))
		msg.Caption = resp
		msg.ReplyMarkup = buttons
		_, err = t.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", err)
		}
	}
}
