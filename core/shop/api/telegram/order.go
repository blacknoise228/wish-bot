package telegram

import (
	"log"
	"strings"
	"wish-bot/core/shop/api/telegram/state"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func (t *Telegram) callbackOrderHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	userMessageData := messageData[query.Message.Chat.ID]

	switch query.Data {
	case "shop_orders":
		go t.deleteLastMessage(chatID)
		t.Service.Order.GetShopOrders(chatID)

	case "admin_orders":
		go t.deleteLastMessage(chatID)
		t.Service.Order.GetAdminOrders(chatID)
	}
	if strings.HasPrefix(query.Data, "cancel_order:") {
		orderID, err := uuid.Parse(strings.TrimPrefix(query.Data, "cancel_order:"))
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении ID заказа!" + err.Error()))
			t.sendMessage(chatID, "Ошибка при получении ID заказа!")
			return
		}
		t.Service.Order.UpdateOrderStatus(chatID, orderID, 5)
	}

	if strings.HasPrefix(query.Data, "aprove_order:") {
		orderID, err := uuid.Parse(strings.TrimPrefix(query.Data, "aprove_order:"))
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении ID заказа!" + err.Error()))
			t.sendMessage(chatID, "Ошибка при получении ID заказа!")
			return
		}
		t.Service.Order.UpdateOrderStatus(chatID, orderID, 2)
	}

	if strings.HasPrefix(query.Data, "paylink_order:") {
		orderID, err := uuid.Parse(strings.TrimPrefix(query.Data, "paylink_order:"))
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении ID заказа!" + err.Error()))
			t.sendMessage(chatID, "Ошибка при получении ID заказа!")
			return
		}
		userMessageData["id"] = orderID.String()
		t.sendMessage(chatID, "Введите ссылку на оплату")
		state.SetUserState(chatID, state.SendPaymentLink)
	}
}

func (t *Telegram) messageOrderHandler(message *tgbotapi.Message, userstate string) {

	chatID := message.Chat.ID

	userMessageData := messageData[message.Chat.ID]

	switch userstate {

	case state.SendPaymentLink:
		orderID, err := uuid.Parse(userMessageData["id"])
		if err != nil {
			log.Println(errornator.CustomError("Ошибка при получении ID заказа!" + err.Error()))
			t.sendMessage(chatID, "Ошибка при получении ID заказа!")
			return
		}
		t.Service.Order.SendPayLink(chatID, orderID, message.Text)
		state.ClearUserState(chatID)
	}
}
