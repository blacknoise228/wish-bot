package service

import (
	"context"
	"fmt"
	"log"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *Service) GetProductsByCategories(chatID int64, categoryID int32) {

	products, err := s.DB.GetProductsByCategory(context.Background(), categoryID)
	if err != nil {
		log.Println(errornator.CustomError("Не удалось получить категорию"))
		return
	}

	for _, product := range products {
		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить в приватные желания",
					"add_wish_private:"+product.ID.String()),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить в публичные желания",
					"add_wish_public:"+product.ID.String()),
			),
		)
		resp := fmt.Sprintf("%v\nОписание: %v\nЦена: %v\nСтатус: %v",
			product.Name, product.Description, product.Price, product.StatusName)
		msg := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(product.Image))
		msg.Caption = resp
		msg.ReplyMarkup = buttons
		_, err = s.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", errornator.CustomError(err.Error()))
		}
	}
}
