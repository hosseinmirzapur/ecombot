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

func GetAPI() *tgbotapi.BotAPI {
	return api
}
