package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleCallback(update tgbotapi.Update) {
	switch update.CallbackData() {
	case "search":
		searchCallback(update)
		return
	case "newest":
		newestCallback(update)
		return
	case "start":
		startCallback(update)
		return
	case "help":
		helpCallback(update)
		return
	default:
		defaultCallback(update)
		return

	}
}

func searchCallback(update tgbotapi.Update) {
	handleBotMessage("search", update.CallbackQuery.Message.Chat.ID)
}

func newestCallback(update tgbotapi.Update) {
	handleBotMessage("newest", update.CallbackQuery.Message.Chat.ID)
}

func startCallback(update tgbotapi.Update) {
	handleBotMessage("start", update.CallbackQuery.Message.Chat.ID)
}

func helpCallback(update tgbotapi.Update) {
	handleBotMessage("help", update.CallbackQuery.Message.Chat.ID)
}

func defaultCallback(update tgbotapi.Update) {
	handleBotMessage("default", update.CallbackQuery.Message.Chat.ID)
}
