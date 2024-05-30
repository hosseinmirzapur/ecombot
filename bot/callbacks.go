package bot

import (
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/hosseinmirzapur/ecombot/database/models"
	"github.com/pocketbase/dbx"
)

func HandleCallback(update tgbotapi.Update, botMode *BotMode) {
	callbackRequest(update)

	switch update.CallbackData() {
	case "/newest":
		newestCallback(update)
		return
	case "/help":
		helpCallback(update)
		return
	default:
		otherCallbacks(update)
		return

	}
}

func callbackRequest(update tgbotapi.Update) {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "در حال بارگذاری...")
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

func otherCallbacks(update tgbotapi.Update) {
	txt := update.CallbackData()

	if strings.Contains(txt, "expand-color") {
		sendExpandColor(update)
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

	if strings.Contains(txt, "video") {
		sendVideos(update)
		return
	}
}

func sendExpandColor(update tgbotapi.Update) {
	colorID := strings.Split(update.CallbackData(), "/")[2]
	chatID := update.CallbackQuery.Message.Chat.ID

	var color models.Color
	err := database.
		DB().
		Select("title", "id").
		From("colors").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": colorID})).
		One(&color)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	showExpandKeyboard(chatID, color)
}

func sendInstaImages(update tgbotapi.Update) {
	colorID := strings.Split(update.CallbackData(), "/")[2]
	chatID := update.CallbackQuery.Message.Chat.ID

	var color models.Color
	err := database.
		DB().
		Select("id", "instagram_images").
		From("colors").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": colorID})).
		One(&color)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	images, err := stringToArray(color.InstagramImages)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(images) == 0 {
		handleBotMessage("هنوز عکس اینستاگرامی برای این محصول ثبت نشده است", chatID)
		return
	}

	for _, image := range images {
		doc, err := getFile(color.ID, image, chatID)
		if err != nil {
			handleErr(err, chatID)
			return
		}
		sendDocToBot(*doc)
		time.Sleep(time.Millisecond * 300)
	}

}

func sendWebImages(update tgbotapi.Update) {
	colorID := strings.Split(update.CallbackData(), "/")[2]
	chatID := update.CallbackQuery.Message.Chat.ID

	var color models.Color
	err := database.
		DB().
		Select("id", "website_images").
		From("colors").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": colorID})).
		One(&color)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	images, err := stringToArray(color.WebsiteImages)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(images) == 0 {
		handleBotMessage("هنوز عکس وبسایتی برای این محصول ثبت نشده است", chatID)
		return
	}

	for _, image := range images {
		doc, err := getFile(color.ID, image, chatID)
		if err != nil {
			handleErr(err, chatID)
			return
		}
		sendDocToBot(*doc)
		time.Sleep(time.Millisecond * 300)
	}
}

func sendVideos(update tgbotapi.Update) {
	colorID := strings.Split(update.CallbackData(), "/")[2]
	chatID := update.CallbackQuery.Message.Chat.ID

	var color models.Color
	err := database.
		DB().
		Select("id", "videos").
		From("colors").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": colorID})).
		One(&color)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	videos, err := stringToArray(color.Videos)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(videos) == 0 {
		handleBotMessage("هنوز ویدئویی برای این محصول ثبت نشده است", chatID)
		return
	}

	for _, video := range videos {
		doc, err := getFile(color.ID, video, chatID)
		if err != nil {
			handleErr(err, chatID)
			return
		}
		sendDocToBot(*doc)
		time.Sleep(time.Millisecond * 300)
	}

}

// func editProduct(update tgbotapi.Update) {
// 	productID := strings.Split(update.CallbackData(), "/")[2]
// 	chatID := update.CallbackQuery.Message.Chat.ID
// 	showEditKeyboard(chatID, productID)
// }

// func deleteProduct(update tgbotapi.Update) {
// 	productID := strings.Split(update.CallbackData(), "/")[2]
// 	chatID := update.CallbackQuery.Message.Chat.ID
// 	showDeleteKeyboard(chatID, productID)
// }
