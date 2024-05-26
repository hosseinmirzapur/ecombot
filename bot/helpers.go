package bot

import (
	"encoding/json"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
)

func stringToArray(urls string) ([]string, error) {
	var stringArray []string
	err := json.Unmarshal([]byte(urls), &stringArray)
	if err != nil {
		return nil, err
	}
	return stringArray, nil
}

func getFile(id, file string, chatID int64) (tgbotapi.DocumentConfig, error) {
	cfg := tgbotapi.DocumentConfig{}
	record, err := database.PB().Dao().FindRecordById("products", id)
	if err != nil {
		return cfg, err
	}

	dataDir := database.PB().DataDir()

	path := fmt.Sprintf("%s/storage/%s/%s", dataDir, record.BaseFilesPath(), file)

	readFile, err := os.Open(path)
	if err != nil {
		return cfg, err
	}

	reader := tgbotapi.FileReader{
		Name:   file,
		Reader: readFile,
	}

	// Create a new Telegram document object
	doc := tgbotapi.NewDocument(chatID, reader)
	doc.Caption = record.GetString("title")

	return doc, nil
}
