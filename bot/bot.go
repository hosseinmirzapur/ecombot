package bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/commands"
	"github.com/hosseinmirzapur/ecombot/messages"
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

func API() *tgbotapi.BotAPI {
	return api
}

func ListenForUpdates() {
	uc := tgbotapi.NewUpdate(0)
	uc.Timeout = 60

	for update := range api.GetUpdatesChan(uc) {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			commands.Handle(update)
		}

		if !update.Message.IsCommand() {
			messages.Handle(update)
		}
	}
}

func HandleErr(err error) {

}
