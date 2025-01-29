package service

import (
	"context"
	"fmt"
	"log"
	db "wish-bot/core/shop/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type OrderService struct {
	db  *db.Queries
	bot *tgbotapi.BotAPI
}

func NewOrderService(db *db.Queries, bot *tgbotapi.BotAPI) Order {
	return OrderService{
		db:  db,
		bot: bot,
	}
}

func (o OrderService) GetAdminOrders(chatID int64) {
	orders, err := o.db.GetOrdersByAdmin(context.Background(), chatID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении заказов!" + err.Error()))
		sendMessage(o.bot, chatID, "Ошибка при получении заказов!")
		return
	}

	for _, order := range orders {
		product, err := o.db.GetProductByID(context.Background(), order.ProductID)
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении заказa!" + err.Error()))
			sendMessage(o.bot, chatID, "Ошибка при получении заказa! ID: "+order.ID.String())
			continue
		}
		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Взять в работу", "aprove_order:"+order.ID.String()),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Нет в наличии", "cancel_order:"+order.ID.String()),
			),
		)

		resp := fmt.Sprintf("ID заказа: %v\nНазвание: %v\nОписание: %v\nЦена: %v\nСтатус товара: %v\nСтатус: %v",
			order.ID, product.Name, product.Description, order.Price, product.StatusName, order.StatusName)

		msg := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(product.Image))
		msg.Caption = resp
		msg.ReplyMarkup = buttons
		_, err = o.bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", err)
		}
	}
}

func (o OrderService) GetShopOrders(chatID int64) {
	admin, err := o.db.GetShopAdminsByAdminID(context.Background(), chatID)
	if err != nil {
		log.Println(errornator.CustomError("Вы не являетесь администратором магазина!" + err.Error()))
		sendMessage(o.bot, chatID, "Вы не являетесь администратором магазина!")
		return
	}

	orders, err := o.db.GetOrdersByShop(context.Background(), admin.ShopID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении заказов!" + err.Error()))
		sendMessage(o.bot, chatID, "Ошибка при получении заказов!")
		return
	}

	for _, order := range orders {
		product, err := o.db.GetProductByID(context.Background(), order.ProductID)
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении заказa!" + err.Error()))
			sendMessage(o.bot, chatID, "Ошибка при получении заказa! ID: "+order.ID.String())
			continue
		}
		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Подтвердить оплату", "aprove_order:"+order.ID.String()),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Отправить ссылку на оплату", "paylink_order:"+order.ID.String()),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Назначить другому администратору", "change_admin_order:"+order.ID.String()),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Нет в наличии", "cancel_order:"+order.ID.String()),
			),
		)

		resp := fmt.Sprintf("ID заказа: %v\nНазвание: %v\nОписание: %v\nЦена: %v\nСтатус товара: %v\nСтатус: %v",
			order.ID, product.Name, product.Description, order.Price, product.StatusName, order.StatusName)

		msg := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(product.Image))
		msg.Caption = resp
		msg.ReplyMarkup = buttons
		_, err = o.bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", errornator.CustomError(err.Error()))
		}
	}
}

func (o OrderService) UpdateOrderStatus(chatID int64, orderID uuid.UUID, status int32) {
	order, err := o.db.GetOrder(context.Background(), orderID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении заказа!" + err.Error()))
		sendMessage(o.bot, chatID, "Ошибка при получении заказа!")
		return
	}

	if order.Status == 2 && status == 4 || order.Status == 2 && status == 5 || order.Status == 2 && status == 1 || order.Status == 3 && status == 4 || order.Status == 3 && status == 5 || order.Status == 3 && status == 1 {
		sendMessage(o.bot, chatID, "Вы не можете изменить статус заказа на этот после оплаты!")
		return
	}

	if err := o.db.UpdateOrderStatus(context.Background(), db.UpdateOrderStatusParams{
		Status:  status,
		ID:      orderID,
		AdminID: chatID,
	}); err != nil {
		log.Println(errornator.CustomError("Ошибка при обновлении статуса заказа!" + err.Error()))
		sendMessage(o.bot, chatID, "Ошибка при обновлении статуса заказа!")
		return
	}
	if status == 2 {
		userInfo, err := o.db.GetUserInfo(context.Background(), order.ConsigneeID)
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении информации о пользователе!" + err.Error()))
			sendMessage(o.bot, chatID, "Ошибка при получении информации о пользователе!")
			return
		}

		resp := fmt.Sprintf("Заказ успешно оплачен!\nДанные для доставки:\nИмя получателя: %v\nТелефон получателя: %v\nАдрес доставки: %v\nКомментарий от получателя:\n%v",
			userInfo.Name, userInfo.Phone, userInfo.Address, userInfo.Description)
		sendMessage(o.bot, order.CustomerID, resp)
	}

	sendMessage(o.bot, chatID, "Статус заказа успешно обновлен!")
}

func (o OrderService) SendPayLink(chatID int64, orderID uuid.UUID, link string) {
	order, err := o.db.GetOrder(context.Background(), orderID)
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении заказа!" + err.Error()))
		sendMessage(o.bot, chatID, "Ошибка при получении заказа!")
		return
	}

	if order.Status == 2 {
		sendMessage(o.bot, chatID, "Заказ уже оплачен!")
		return
	}
	sendMessage(o.bot, order.CustomerID, "Ссылка на оплату: "+link)
	sendMessage(o.bot, chatID, "Ссылка на оплату успешно отправлена!")
}
