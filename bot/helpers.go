package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func getFile(id, file string, chatID int64) (*tgbotapi.DocumentConfig, error) {
	record, err := database.PB().Dao().FindRecordById("products", id)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/%s", os.Getenv("FILES_BASEURL"), record.BaseFilesPath(), file)

	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	// Check for errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download failed: %s", resp.Status)
	}

	// Create a temporary file
	createdFile, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	// Download the file to the temporary file
	_, err = io.Copy(createdFile, resp.Body)
	if err != nil {
		return nil, err
	}

	reader := tgbotapi.FileReader{
		Name:   file,
		Reader: createdFile, // Use temporary file reader
	}

	// Create a new Telegram document object
	doc := tgbotapi.NewDocument(chatID, reader)
	doc.Caption = record.GetString("title")

	return &doc, nil
}
