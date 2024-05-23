package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/hosseinmirzapur/ecombot/database/models"
)

func HandleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		startCommand(update)
		return
	case "newest":
		newestCommand(update)
		return
	case "search":
		searchCommand(update)
		return
	case "help":
		helpCommand(update)
		return
	default:
		defaultCommand(update)
		return
	}
}

func startCommand(update tgbotapi.Update) {
	tgID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	var user models.User
	err := database.DB().Where("telegram_id = ?", tgID).Find(&user).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if user.ID == 0 {
		register(tgID, chatID)
		return
	}

	homeInlineKeyboard(chatID)

}

func register(tgID int64, chatID int64) {
	var user models.User

	user.TelegramID = tgID
	err := database.DB().Create(&user).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	handleBotMessage("Your account has been registered successfully!", chatID)
}

func newestCommand(update tgbotapi.Update) {
	var products []models.Product
	chatID := update.Message.Chat.ID

	err := database.DB().Select("title").Find(&products).Order("id DESC").Take(10).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("back", "/start"),
		),
	)

	for _, product := range products {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(product.Title, product.Title),
		))

	}

	msg := tgbotapi.NewMessage(chatID, "newest products")
	msg.ReplyMarkup = keyboard
	sendToBot(msg)
}

func searchCommand(update tgbotapi.Update) {}

func helpCommand(update tgbotapi.Update) {}

func defaultCommand(update tgbotapi.Update) {}
