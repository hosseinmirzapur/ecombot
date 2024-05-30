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

func singleProductInlineKeyboard(product models.Product, colors []models.Color, chatID int64, botMode *BotMode) {
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

	if len(colors) > 0 {
		var rows [][]tgbotapi.InlineKeyboardButton

		for _, color := range colors {
			rows = append(rows, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("رنگ %s", color.Title), fmt.Sprintf("/expand-color/%s", color.ID)),
			))
		}

		if botMode.IsAdminMode() {
			rows = append(rows, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("ویرایش محصول", fmt.Sprintf("/edit/%s", product.ID)),
				tgbotapi.NewInlineKeyboardButtonData("حذف محصول", fmt.Sprintf("/delete/%s", product.ID)),
			))
		}
		keyboard := tgbotapi.NewInlineKeyboardMarkup(rows...)

		msg.ReplyMarkup = keyboard
	}

	msg.ParseMode = tgbotapi.ModeMarkdown

	sendToBot(msg)

}

func showExpandKeyboard(chatID int64, color models.Color) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("عکس های اینستاگرام", fmt.Sprintf("/insta/%s", color.ID)),
			tgbotapi.NewInlineKeyboardButtonData("عکس های وبسایت", fmt.Sprintf("/web/%s", color.ID)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("ویدئو ها", fmt.Sprintf("/video/%s", color.ID)),
		),
	)

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("اطلاعات رنگ %s این محصول", color.Title))
	msg.ReplyMarkup = keyboard

	sendToBot(msg)
}

// func showEditKeyboard(chatID int64, productID string) {
// 	keyboard := tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("عنوان", fmt.Sprintf("/edit-title/%s", productID)),
// 			tgbotapi.NewInlineKeyboardButtonData("توضیحات", fmt.Sprintf("/edit-description/%s", productID)),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("قیمت", fmt.Sprintf("/edit-price/%s", productID)),
// 			tgbotapi.NewInlineKeyboardButtonData("کد محصول", fmt.Sprintf("/edit-code/%s", productID)),
// 		),
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("لیست عکس های اینستاگرامی", fmt.Sprintf("/edit-insta/%s", productID)),
// 			tgbotapi.NewInlineKeyboardButtonData("لیست عکس های وبسایت", fmt.Sprintf("/edit-web/%s", productID)),
// 		),
// 	)

// 	msg := tgbotapi.NewMessage(chatID, "کدام مورد را میخواهید ویرایش کنید؟")
// 	msg.ReplyMarkup = keyboard
// 	sendToBot(msg)
// }

// func showDeleteKeyboard(chatID int64, productID string) {
// 	keyboard := tgbotapi.NewInlineKeyboardMarkup(
// 		tgbotapi.NewInlineKeyboardRow(
// 			tgbotapi.NewInlineKeyboardButtonData("تایید میکنم", fmt.Sprintf("/accept-delete/%s", productID)),
// 			tgbotapi.NewInlineKeyboardButtonData("منصرف شدم", fmt.Sprintf("/reject-delete/%s", productID)),
// 		),
// 	)

// 	msg := tgbotapi.NewMessage(chatID, "آیا از حذف این محصول مطمئن هستید؟")
// 	msg.ReplyMarkup = keyboard
// 	sendToBot(msg)
// }
