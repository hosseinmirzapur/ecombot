package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleCallback(update tgbotapi.Update) {
	callbackRequest(update)

	switch update.CallbackData() {
	case "/newest":
		newestCallback(update)
		return
	case "/help":
		helpCallback(update)
		return
	default:
		defaultCallback(update)
		return

	}
}

func callbackRequest(update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		// do nothing
		return
	}
}

func newestCallback(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	newestCommand(chatID)
}

func helpCallback(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	helpCommand(chatID)
}

func defaultCallback(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	defaultCommand(chatID)
}
