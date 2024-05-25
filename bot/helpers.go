package bot

import (
	"encoding/json"
	"fmt"
	"os"

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

func fileURL(id, file string) (string, error) {
	record, err := database.PB().Dao().FindRecordById("products", id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s", os.Getenv("FILES_BASEURL"), record.BaseFilesPath(), file), nil
}
