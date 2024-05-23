package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func homeInlineKeyboard(chatID int64) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("جدید ترین ها", "newest"),
			tgbotapi.NewInlineKeyboardButtonData("جستجو", "search"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("راهنمایی", "help"),
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
