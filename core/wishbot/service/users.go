package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "wish-bot/core/wishbot/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func (t *Service) StartMessageHandler(message *tgbotapi.Message) {
	t.sendMessage(message.Chat.ID, StartMessage)
	buttonMessage := "Добро пожаловать в Wish Bot!"
	button := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Создать профиль", "register"),
		),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, buttonMessage)
	msg.ReplyMarkup = button

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке сообщения:", err)
	}
}
func (t *Service) CreateUserHandler(ctx context.Context, message *tgbotapi.Message) error {
	_, err := t.DB.GetUser(ctx, message.Chat.ID)
	if err == nil {
		if sendErr := t.sendMessage(message.Chat.ID, "Вы уже зарегистрированы!"); sendErr != nil {
			return sendErr
		}
		return nil
	}

	_, err = t.DB.GetUserByUsername(ctx, message.Text)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return errornator.CustomError(err.Error())
		}
	} else {
		if sendErr := t.sendMessage(message.Chat.ID, "Такой username занят! Придумайте другой и попробуйте снова."); sendErr != nil {
			return sendErr
		}
		return nil
	}
	_, err = t.DB.CreateUser(ctx, db.CreateUserParams{
		ChatID:   message.Chat.ID,
		Username: message.Text,
	})
	if err != nil {
		return err
	}
	if err := t.sendMessage(message.Chat.ID, "Пользователь успешно зарегистрирован! Ваш ник: "+message.Text+"."); err != nil {
		return err
	}
	return nil
}

func (t *Service) UpdateUserHandler(ctx context.Context, message *tgbotapi.Message) error {

	usr, err := t.DB.UpdateUser(ctx, db.UpdateUserParams{
		ChatID:   message.Chat.ID,
		Username: message.Text,
	})
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Успешно!\nВаш новый username: \n%v", usr.Username)
	t.sendMessage(message.Chat.ID, msg)

	return nil
}

func (t *Service) DeleteUserHandler(query *tgbotapi.CallbackQuery) error {
	friends, err := t.DB.GetPendingFriendships(context.Background(), query.Message.Chat.ID)
	if err != nil {
		return err
	}
	if len(friends) != 0 {
		for _, friend := range friends {
			t.DB.DeleteFriendship(context.Background(),
				db.DeleteFriendshipParams{ChatID: query.Message.Chat.ID, FriendID: friend.FriendID})

			t.DB.DeleteFriendship(context.Background(),
				db.DeleteFriendshipParams{ChatID: friend.FriendID, FriendID: query.Message.Chat.ID})
		}
	}
	wishes, err := t.DB.GetWishesForUser(context.Background(), query.Message.Chat.ID)
	if err != nil {
		return err
	}
	if len(wishes) != 0 {
		for _, wish := range wishes {
			t.DB.DeleteWish(context.Background(), db.DeleteWishParams{ChatID: wish.ChatID, ID: wish.ID})
		}
	}
	return t.DB.DeleteUser(context.Background(), query.Message.Chat.ID)
}
