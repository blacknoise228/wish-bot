package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	db "wish-bot/core/wishbot/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (t *Service) CreateWish(chatID int64, productID uuid.UUID, status int32) {
	_, err := t.DB.CreateWish(context.Background(), db.CreateWishParams{
		ChatID:    chatID,
		ProductID: productID,
		Status:    status,
	})
	if err != nil {
		t.sendMessage(chatID, "Ошибка при добавлении желания!")
		log.Println(errornator.CustomError(err.Error()))
		return
	}
	t.sendMessage(chatID, "Желание добавлено!")
}

func (t *Service) GetMyWishes(chatID int64) {
	wishes, err := t.DB.GetWishesForUser(context.Background(), chatID)
	if err != nil {
		t.sendMessage(chatID, "Ошибка при получении желаний")
		log.Println(errornator.CustomError(err.Error()))
		return
	}
	if len(wishes) == 0 {
		t.sendMessage(chatID, "У вас нет желаний!")
		return
	}

	for _, wish := range wishes {

		product, err := t.DB.GetProductByID(context.Background(), wish.ProductID)
		if err != nil {
			log.Println(errornator.CustomError(err.Error()))
		}

		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Редактировать статус", "update_wish:"+strconv.Itoa(int(wish.ID))),
				tgbotapi.NewInlineKeyboardButtonData("Удалить желание", "remove_wish:"+strconv.Itoa(int(wish.ID))),
			),
		)
		resp := fmt.Sprintf("Имя пользователя: %v\nОписание: %v\nЦена: %v\nСтатус: %v",
			wish.Username, product.Description, product.Price, wish.StatusName)
		msg := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(product.Image))
		msg.Caption = resp
		msg.ReplyMarkup = buttons
		_, err = t.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", errornator.CustomError(err.Error()))
		}
	}
}

func (t *Service) GetUserWishes(chatID int64, friendUsername string) {
	friend, err := t.DB.GetUserByUsername(context.Background(), friendUsername)
	if err != nil {
		t.sendMessage(chatID, "Такого пользователя не существует!")
		return
	}
	wishes, err := t.DB.GetWishesPublic(context.Background(), db.GetWishesPublicParams{
		ChatID:   friend.ChatID,
		ChatID_2: chatID,
	})
	if err != nil {
		log.Println(errornator.CustomError("Ошибка при получении желаний!" + err.Error()))
		t.sendMessage(chatID, "Ошибка при получении желаний")
		return
	}
	if len(wishes) == 0 {
		t.sendMessage(chatID, "У пользователя нет желаний!")
		return
	}
	for _, wish := range wishes {
		product, err := t.DB.GetProductByID(context.Background(), wish.ProductID)
		if err != nil {
			log.Println(errornator.CustomError("Такого товара не существует!" + err.Error()))
		}

		buttons := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Подарить", "create_order:"+strconv.Itoa(int(wish.ID))),
			),
		)
		resp := fmt.Sprintf("Имя пользователя: %v\nОписание: %v\nЦена: %v\nСтатус: %v",
			wish.Username, product.Description, product.Price, wish.StatusName)
		msg := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(product.Image))
		msg.Caption = resp
		msg.ReplyMarkup = buttons
		_, err = t.Bot.Send(msg)
		if err != nil {
			log.Println("Ошибка при отправке встроенного меню:", errornator.CustomError(err.Error()))
		}
	}
}

func (t *Service) DeleteWish(chatID int64, wishID string) {

	wish, err := strconv.Atoi(wishID)
	if err != nil {
		t.sendMessage(chatID, "Некорректное значение ID")
		return
	}

	err = t.DB.DeleteWish(context.Background(), db.DeleteWishParams{ChatID: chatID, ID: int32(wish)})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			t.sendMessage(chatID, "Такого желания не существует!")
			return
		}

		t.sendMessage(chatID, "Ошибка при удалении желания!")
		return
	}

	t.sendMessage(chatID, "Желание удалено!")
}

func (t *Service) UpdateWish(chatID int64, wishID string) {
	var stat int32
	wishId, err := strconv.Atoi(wishID)
	if err != nil {
		log.Println(chatID, "Некорректное значение ID")
		return
	}
	status, err := t.DB.GetWishByID(context.Background(), int32(wishId))
	if err != nil {
		log.Println(chatID, "Такого желания не существует!")
		return
	}

	if status.Status == 1 {
		stat = 2
	} else {
		stat = 1
	}

	_, err = t.DB.UpdateWishStatus(context.Background(), db.UpdateWishStatusParams{
		ID:     int32(wishId),
		ChatID: chatID,
		Status: stat,
	})
	if err != nil {
		log.Println(chatID, "Ошибка при обновлении статуса желания!")
		return
	}

	t.sendMessage(chatID, "Статус желания успешно обновлен!")
}
