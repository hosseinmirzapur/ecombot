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

func SetDebug(flag bool) {
	bot.Debug = flag
}

func ListenForUpdates() {
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = 60

	for update := range bot.GetUpdatesChan(uc) {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			HandleCommand(update)
		}

		if !update.Message.IsCommand() {
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

}
