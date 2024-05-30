package bot

import (
	"encoding/json"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
)

func stringToArray(strs string) ([]string, error) {
	var stringArray []string
	err := json.Unmarshal([]byte(strs), &stringArray)
	if err != nil {
		return nil, err
	}
	return stringArray, nil
}

func getFile(id, file string, chatID int64) (*tgbotapi.DocumentConfig, error) {
	cfg := tgbotapi.DocumentConfig{}
	record, err := database.PB().Dao().FindRecordById("colors", id)
	if err != nil {
		return nil, err
	}

	dataDir := database.PB().DataDir()

	path := fmt.Sprintf("%s/storage/%s/%s", dataDir, record.BaseFilesPath(), file)

	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := tgbotapi.FileReader{
		Name:   file,
		Reader: readFile,
	}

	// Create a new Telegram document object
	cfg = tgbotapi.NewDocument(chatID, reader)
	cfg.Caption = record.GetString("title")

	// remove created file
	err = os.Remove(path)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
