package telegram

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"wish-bot/core/shop/api/telegram/state"
	"wish-bot/pkg/errornator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackProductHandler(query *tgbotapi.CallbackQuery) {

	chatID := query.Message.Chat.ID

	userMessageData := messageData[query.Message.Chat.ID]

	switch query.Data {
	case "add_product":
		go t.deleteLastMessage(chatID)
		t.sendMessage(chatID, "Добавьте название")
		state.SetUserState(chatID, state.AddProductName)

	}
	if strings.HasPrefix(query.Data, "update_product:") {
		userMessageData["id"] = strings.TrimPrefix(query.Data, "update_product:")
	}
}

func (t *Telegram) createProductHandler(message *tgbotapi.Message, userstate string) {

	chatID := message.Chat.ID

	userMessageData := messageData[message.Chat.ID]

	switch userstate {

	case state.AddProductName:
		userMessageData["name"] = message.Text
		t.sendMessage(chatID, "Добавьте описание")
		state.SetUserState(chatID, state.AddProductDesc)

	case state.AddProductDesc:
		userMessageData["desc"] = message.Text
		t.sendMessage(chatID, "Выберите категорию")
		t.productSelectCategory(chatID)
		state.SetUserState(chatID, state.AddProductCategory)

	case state.AddProductCategory:
		userMessageData["category"] = message.Text

		msg := tgbotapi.NewMessage(chatID, "Добавьте цену")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // true = убрать клавиатуру для всех пользователей
		t.Bot.Send(msg)

		state.SetUserState(chatID, state.AddProductPrice)

	case state.AddProductPrice:
		userMessageData["price"] = message.Text
		t.sendMessage(chatID, "Добавьте изображение")
		state.SetUserState(chatID, state.AddProductImage)

	case state.AddProductImage:
		photoPath := t.downloadPhoto(message)
		userMessageData["image"] = photoPath
		t.Service.Product.CreateProduct(chatID, userMessageData)
		t.sendInlineMenu(chatID)
		t.sendMenuButton(chatID)
		state.ClearUserState(chatID)
	}
}

func (t *Telegram) updateProductHandler(message *tgbotapi.Message, userstate string) {

	chatID := message.Chat.ID

	userMessageData := messageData[message.Chat.ID]

	switch userstate {

	case state.UpdateProductStatus:
		t.productSelectStatus(chatID)
		t.Service.Product.UpdateProductStatus(chatID, userMessageData["id"], message.Text)
		state.ClearUserState(chatID)

	case state.UpdateProductName:
		userMessageData["name"] = message.Text
		t.sendMessage(chatID, "Добавьте описание")
		state.SetUserState(chatID, state.UpdateProductDesc)

	case state.UpdateProductDesc:
		userMessageData["desc"] = message.Text
		t.sendMessage(chatID, "Выберите категорию")
		t.productSelectCategory(chatID)
		state.SetUserState(chatID, state.UpdateProductCategory)

	case state.UpdateProductCategory:
		userMessageData["category"] = message.Text

		msg := tgbotapi.NewMessage(chatID, "Добавьте цену")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // true = убрать клавиатуру для всех пользователей
		t.Bot.Send(msg)

		state.SetUserState(chatID, state.UpdateProductPrice)

	case state.UpdateProductPrice:
		userMessageData["price"] = message.Text
		t.sendMessage(chatID, "Добавьте изображение")
		state.SetUserState(chatID, state.UpdateProductImage)

	case state.UpdateProductImage:
		photoPath := t.downloadPhoto(message)
		userMessageData["image"] = photoPath
		t.Service.Product.CreateProduct(chatID, userMessageData)
		t.sendInlineMenu(chatID)
		t.sendMenuButton(chatID)
		state.ClearUserState(chatID)
	}
}

func (t *Telegram) productSelectCategory(chatID int64) {
	buttons := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Цветы"),
		tgbotapi.NewKeyboardButton("Одежда"),
		tgbotapi.NewKeyboardButton("Электроника"),
		tgbotapi.NewKeyboardButton("Игрушки"),
		tgbotapi.NewKeyboardButton("Аксессуары"),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите категорию")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
	}
}

func (t *Telegram) productSelectStatus(chatID int64) {
	buttons := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("В наличии"),
		tgbotapi.NewKeyboardButton("Нет в наличии"),
	)
	msg := tgbotapi.NewMessage(chatID, "Выберите статус")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

	if _, err := t.Bot.Send(msg); err != nil {
		log.Println("Ошибка при отправке встроенного меню:", err)
		errornator.CustomError(err.Error())
	}
}
func (t *Telegram) downloadPhoto(message *tgbotapi.Message) string {
	if len(message.Photo) == 0 {
		return ""
	}

	photo := message.Photo[len(message.Photo)-1]
	fileID := photo.FileID

	fileConfig := tgbotapi.FileConfig{FileID: fileID}
	file, err := t.Bot.GetFile(fileConfig)
	if err != nil {
		log.Println("Ошибка при получении информации о файле:", err)
		return ""
	}

	downloadURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", t.Bot.Token, file.FilePath)

	// Шаг 4: Скачиваем файл
	resp, err := http.Get(downloadURL)
	if err != nil {
		log.Println("Ошибка при скачивании:", err)
		return ""
	}
	defer resp.Body.Close()

	// Шаг 5: Сохраняем файл локально
	// Можно придумать свою логику формирования имени файла: UUID, метка времени и т.д.
	localFileName := fmt.Sprintf("%s/%s.jpg", t.Config.App.Photos, time.Now().Format("20060102_150405"))
	out, err := os.Create(localFileName)
	if err != nil {
		log.Println("Ошибка при создании файла:", err)
		return ""
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Ошибка при записи файла:", err)
		return ""
	}
	return localFileName
}
