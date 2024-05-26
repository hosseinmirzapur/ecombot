package bot

import (
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/hosseinmirzapur/ecombot/database/models"
	"github.com/pocketbase/dbx"
)

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
	txt := update.CallbackData()

	if !strings.Contains(txt, "insta") &&
		!strings.Contains(txt, "web") &&
		!strings.Contains(txt, "videos") {
		defaultCommand(chatID)
		return
	}

	if strings.Contains(txt, "insta") {
		sendInstaImages(update)
		return
	}

	if strings.Contains(txt, "web") {
		sendWebImages(update)
		return
	}

	if strings.Contains(txt, "videos") {
		sendVideos(update)
		return
	}
}

func sendInstaImages(update tgbotapi.Update) {
	code := strings.Split(update.CallbackData(), "/")[4]
	chatID := update.CallbackQuery.Message.Chat.ID

	var products []models.Product
	err := database.
		DB().
		Select("instagram_images", "id").
		From("products").
		Where(dbx.NewExp("code = {:code}", dbx.Params{"code": code})).
		All(&products)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	images, err := stringToArray(products[0].InstagramImages)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(images) == 0 {
		handleBotMessage("هنوز عکس اینستاگرامی برای این محصول ثبت نشده است", chatID)
		return
	}

	for _, image := range images {
		doc, err := getFile(products[0].ID, image, chatID)
		if err != nil {
			handleErr(err, chatID)
			return
		}
		sendDocToBot(doc)
		time.Sleep(time.Millisecond * 300)
	}

}

func sendWebImages(update tgbotapi.Update) {
	code := strings.Split(update.CallbackData(), "/")[4]
	chatID := update.CallbackQuery.Message.Chat.ID

	var products []models.Product
	err := database.
		DB().
		Select("website_images", "id").
		From("products").
		Where(dbx.NewExp("code = {:code}", dbx.Params{"code": code})).
		All(&products)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	images, err := stringToArray(products[0].WebsiteImages)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(images) == 0 {
		handleBotMessage("هنوز عکس وبسایتی برای این محصول ثبت نشده است", chatID)
		return
	}

	for _, image := range images {
		doc, err := getFile(products[0].ID, image, chatID)
		if err != nil {
			handleErr(err, chatID)
			return
		}
		sendDocToBot(doc)
		time.Sleep(time.Millisecond * 300)
	}
}

func sendVideos(update tgbotapi.Update) {
	code := strings.Split(update.CallbackData(), "/")[3]
	chatID := update.CallbackQuery.Message.Chat.ID

	var products []models.Product
	err := database.
		DB().
		Select("videos", "id").
		From("products").
		Where(dbx.NewExp("code = {:code}", dbx.Params{"code": code})).
		All(&products)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	videos, err := stringToArray(products[0].Videos)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(videos) == 0 {
		handleBotMessage("هنوز ویدئویی برای این محصول ثبت نشده است", chatID)
		return
	}

	for _, video := range videos {
		doc, err := getFile(products[0].ID, video, chatID)
		if err != nil {
			handleErr(err, chatID)
			return
		}
		sendDocToBot(doc)
		time.Sleep(time.Millisecond * 300)
	}

}
