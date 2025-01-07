package telegram

import (
	"context"
	"errors"
	"fmt"
	"log"
	db "wish-bot/internal/db/sqlc"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func (t *Telegram) startMessageHandler(message *tgbotapi.Message) {
	t.sendMessage(message.Chat.ID, startMessage)
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
func (t *Telegram) createUserHandler(ctx context.Context, message *tgbotapi.Message) error {
	_, err := t.services.User.GetUser(ctx, message.Chat.ID)
	if err == nil {
		if sendErr := t.sendMessage(message.Chat.ID, "Вы уже зарегистрированы!"); sendErr != nil {
			return sendErr
		}
		return nil
	}

	_, err = t.services.User.GetUserByUsername(ctx, message.Text)
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
	_, err = t.services.User.CreateUser(ctx, db.CreateUserParams{
		ChatID:   message.Chat.ID,
		Username: message.Text,
	})
	if err != nil {
		return err
	}
	if err := t.sendMessage(message.Chat.ID, "Пользователь успешно зарегистрирован!"); err != nil {
		return err
	}
	return nil
}

func (t *Telegram) getUserHandler(ctx context.Context, message *tgbotapi.Message) (*db.User, error) {
	user, err := t.services.User.GetUser(ctx, message.Chat.ID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (t *Telegram) updateUserHandler(ctx context.Context, message *tgbotapi.Message) error {

	usr, err := t.services.User.UpdateUser(ctx, db.UpdateUserParams{
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

func (t *Telegram) deleteUserHandler(query *tgbotapi.CallbackQuery) error {
	friends, err := t.services.Friend.GetPendingFriendships(context.Background(), query.Message.Chat.ID)
	if err != nil {
		return err
	}
	if len(friends) != 0 {
		for _, friend := range friends {
			if friend.ChatID != query.Message.Chat.ID {
				t.services.Friend.DeleteFriendship(context.Background(), friend.ChatID, query.Message.Chat.ID)
			}
			t.services.Friend.DeleteFriendship(context.Background(), query.Message.Chat.ID, friend.FriendID)
		}
	}
	return t.services.User.DeleteUser(context.Background(), query.Message.Chat.ID)
}
