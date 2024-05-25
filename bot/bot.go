package bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func RegisterBot() error {
	// instantiate bit
	b, err := tgbotapi.NewBotAPI(
		os.Getenv("BOT_TOKEN"),
	)
	if err != nil {
		return err
	}

	// remove any redundant webhooks
	if _, err = b.Request(tgbotapi.DeleteWebhookConfig{}); err != nil {
		return err
	}

	// set global api
	bot = b
	return nil
}

func RegisterCommands() error {
	commands := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     "/start",
			Description: "شروع به کار ربات",
		},
		tgbotapi.BotCommand{
			Command:     "/newest",
			Description: "نمایش جدیدترین محصولات فروشگاه ما",
		},
		tgbotapi.BotCommand{
			Command:     "/search",
			Description: "جستجو در بین کالا های فروشگاه ما",
		},
		tgbotapi.BotCommand{
			Command:     "/help",
			Description: "راهنمای استفاده از ربات",
		},
	)
	if _, err := bot.Request(commands); err != nil {
		return err
	}

	return nil
}

func SetDebug(flag bool) {
	bot.Debug = flag
}

func ListenForUpdates() {
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = 60

	for update := range bot.GetUpdatesChan(uc) {
		if update.CallbackData() != "" {
			HandleCallback(update)
		} else if update.Message == nil {
			continue
		} else if update.Message.IsCommand() {
			HandleCommand(update)
		} else {
			HandleMessage(update)

		}

	}
}

func handleErr(err error, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, err.Error())
	bot.Send(msg)
}

func handleBotMessage(msg string, chatID int64) {
	tgMessage := tgbotapi.NewMessage(chatID, msg)
	_, err := bot.Send(tgMessage)
	if err != nil {
		handleErr(err, chatID)
		return
	}
}

func sendToBot(msg tgbotapi.MessageConfig) {
	_, err := bot.Send(msg)
	if err != nil {
		handleErr(err, msg.ChatID)
		return
	}
}

func sendImageToBot(image tgbotapi.PhotoConfig) {
	_, err := bot.Send(image)
	if err != nil {
		handleErr(err, image.ChatID)
		return
	}
}

func sendVideoToBot(video tgbotapi.VideoConfig) {
	_, err := bot.Send(video)
	if err != nil {
		handleErr(err, video.ChatID)
		return
	}
}
