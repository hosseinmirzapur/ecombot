package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/hosseinmirzapur/ecombot/database/models"
)

func HandleCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		start(update)
		return
	case "newest":
		newest(update)
		return
	case "search":
		search(update)
		return
	case "help":
		help(update)
		return
	default:
		handleDefault(update)
	}
}

func start(update tgbotapi.Update) {
	tgID := update.Message.From.ID
	chatID := update.Message.Chat.ID

	var user models.User
	err := database.FindOne(user, fmt.Sprintf("TelegramID = %d", tgID)).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if user.ID == 0 {
		register(tgID, chatID)
		return
	}

}

func register(tgID int64, chatID int64) {
	err := database.Create(models.User{TelegramID: tgID}).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	handleBotMessage("Your account has been registered successfully!", chatID)
}

func newest(update tgbotapi.Update) {
	var products []models.Product
	chatID := update.Message.Chat.ID

	err := database.DB().Select("Title").Find(&products).Order("ID DESC").Take(10).Error
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

func search(update tgbotapi.Update) {}

func help(update tgbotapi.Update) {}

func handleDefault(update tgbotapi.Update) {}
