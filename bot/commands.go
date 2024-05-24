package bot

import (
	"fmt"
	"strings"

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

	err := database.DB().Select("title").Limit(10).Order("id DESC").Find(&products).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(products) == 0 {
		handleBotMessage("محصولی در فروشگاه موجود نیست", chatID)
		return
	}

	productsKeyboard(products, chatID)
}

func searchCommand(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	// extract search query after `/search` syntax
	updateMsg := update.Message.Text
	splitMsg := strings.Split(updateMsg, " ")
	if len(splitMsg) < 2 {
		handleErr(fmt.Errorf(`
		برای جستجو در بین محصولات با فرمت زیر جستجو نمایید:

		/search نام محصول

		ابتدا عبارت search/ را قرار دهید و سپس نام محصول مورد نظر خود را وارد کنید تا جستجو صورت پذیرد
		`), chatID)
		return
	}
	searchQ := strings.Join(splitMsg[1:], ">")

	var products []models.Product
	err := database.DB().Where("title LIKE ?", "%"+searchQ+"%").Find(&products).Error
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(products) == 0 {
		handleBotMessage("پس از جستجو موردی جهت نمایش یافت نشد", chatID)
		return
	}

	productsKeyboard(products, chatID)

}

func helpCommand(update tgbotapi.Update) {}

func defaultCommand(update tgbotapi.Update) {}
