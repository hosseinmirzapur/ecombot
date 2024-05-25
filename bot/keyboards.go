package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database/models"
)

func homeInlineKeyboard(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("جدید ترین ها", "/newest"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("راهنما", "/help"),
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

func singleProductInlineKeyboard(product models.Product, chatID int64) {
	// show the metadata
	txt := `
	**%s**

	عکس و ویدئو ها با محصولی که به دست شما میرسد کاملا مطابقت دارد

	**مشخصات**:
	%s

	**قیمت**: %s
	**کد**: %s (برای راحت پیدا کردن محصول در وبسایت یا تلگرام)
	`

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(txt, product.Title, product.Description, product.Price, product.Code))

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("عکس برای اینستاگرام", fmt.Sprintf("/insta/images/product/%s", product.Code)),
			tgbotapi.NewInlineKeyboardButtonData("عکس برای وبسایت", fmt.Sprintf("/web/images/product/%s", product.Code)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ویدئو های محصول", fmt.Sprintf("/videos/product/%s", product.Code)),
		),
	)

	msg.ReplyMarkup = keyboard
	msg.ParseMode = tgbotapi.ModeMarkdown

	sendToBot(msg)

}
