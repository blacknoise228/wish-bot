package telegram

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

func (t *Telegram) callbackOrderHandler(query *tgbotapi.CallbackQuery) {
	if strings.HasPrefix(query.Data, "create_order:") {
		strID := strings.TrimPrefix(query.Data, "create_order:")
		id, _ := strconv.Atoi(strID)
		t.Service.CreateOrder(query.Message.Chat.ID, query.Message.From.UserName, int32(id))
	}
	if strings.HasPrefix(query.Data, "cancel_order:") {
		strID := strings.TrimPrefix(query.Data, "cancel_order:")
		id := uuid.MustParse(strID)
		t.Service.UpdateOrderStatus(query.Message.Chat.ID, id, 4)
	}
}
