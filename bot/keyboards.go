package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database/models"
)

func homeInlineKeyboard(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("جدید ترین ها", "newest"),
			tgbotapi.NewInlineKeyboardButtonData("جستجو", "search"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("راهنما", "help"),
			tgbotapi.NewInlineKeyboardButtonData("راه اندازی مجدد", "start"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, `
	به ربات تلگرامی ما خوش آمدید
	از منوی زیر میتوانید درخواست خود را وارد نمایید
	`)
	msg.ReplyMarkup = keyboard

	sendToBot(msg)

}

func productsKeyboard(products []models.Product, chatID int64) {

	var keyboard [][]tgbotapi.KeyboardButton

	for _, product := range products {
		keyboard = append(keyboard, tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(product.Title),
		))
	}

	msg := tgbotapi.NewMessage(chatID, "لیست محصولات")
	msg.ReplyMarkup = tgbotapi.NewOneTimeReplyKeyboard(keyboard...)
	sendToBot(msg)
}
