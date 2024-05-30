package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/hosseinmirzapur/ecombot/database/models"
	"github.com/pocketbase/dbx"
)

func HandleMessage(update tgbotapi.Update, botMode *BotMode) {
	chatID := update.Message.Chat.ID
	updateTxt := update.Message.Text

	var products []models.Product
	err := database.
		DB().
		Select("id", "title", "code", "description", "created", "colors").
		From("products").
		Where(dbx.Like("title", updateTxt)).
		OrWhere(dbx.NewExp("title = {:title}", dbx.Params{"title": updateTxt})).
		OrWhere(dbx.NewExp("code = {:code}", dbx.Params{"code": updateTxt})).
		OrderBy("created DESC").
		All(&products)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(products) == 0 {
		handleBotMessage(`
		محصولی با مشخصات وارد شده یافت نشد.
		یا کد محصول را وارد نمایید یا نام تقریبی محصول را بنویسید و یا از طریق دستور سرچ ربات محصول مورد نظر خود را پیدا کنید.
		`, chatID)
		return
	}

	if len(products) > 1 {
		productsKeyboard(products, chatID)
		return
	}

	colorIDs, err := stringToArray(products[0].Colors)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	var colors []models.Color
	if len(colorIDs) > 0 {
		for _, colorID := range colorIDs {
			var color models.Color
			err := database.DB().Select("id", "title").From("colors").Where(dbx.NewExp("id = {:id}", dbx.Params{"id": colorID})).One(&color)
			if err != nil {
				handleErr(err, chatID)
				return
			}
			colors = append(colors, color)
		}
	}

	singleProductInlineKeyboard(products[0], colors, chatID, botMode)
}
