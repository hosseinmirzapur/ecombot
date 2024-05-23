package bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var api *tgbotapi.BotAPI

func RegisterBot() error {
	// instantiate bit
	bot, err := tgbotapi.NewBotAPI(
		os.Getenv("BOT_TOKEN"),
	)
	if err != nil {
		return err
	}

	// set global api
	api = bot
	return nil
}

func SetDebug(flag bool) {
	api.Debug = flag
}

func ListenForUpdates() {
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = 60

	for update := range api.GetUpdatesChan(uc) {
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

}

func handleBotMessage(msg interface{}, chatID int64) {

}

func sendToBot(msg tgbotapi.MessageConfig) {

}
