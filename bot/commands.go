package bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/hosseinmirzapur/ecombot/database/models"
	"github.com/pocketbase/dbx"
)

func HandleCommand(update tgbotapi.Update, botMode *BotMode) {
	chatID := update.Message.Chat.ID
	switch update.Message.Command() {
	case "start":
		startCommand(chatID)
		return
	case "newest":
		newestCommand(chatID)
		return
	case "search":
		searchCommand(update, chatID)
		return
	case "help":
		helpCommand(chatID)
		return
	default:
		defaultCommand(chatID)
		return
	}
}

func startCommand(chatID int64) {
	homeInlineKeyboard(chatID)
}

func newestCommand(chatID int64) {
	var products []models.Product

	err := database.
		DB().
		Select("title").
		From("products").
		OrderBy("created").
		Limit(10).
		All(&products)

	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(products) == 0 {
		handleBotMessage("محصولی در فروشگاه موجود نیست", chatID)
		return
	}

	productsKeyboard(products, chatID)
}

func searchCommand(update tgbotapi.Update, chatID int64) {

	// extract search query after `/search` syntax
	updateMsg := update.Message.Text
	splitMsg := strings.Split(updateMsg, " ")
	if len(splitMsg) < 2 {
		handleErr(fmt.Errorf(`
		برای جستجو در بین محصولات با فرمت زیر جستجو نمایید:

		/search نام محصول

		ابتدا عبارت search/ را قرار دهید و سپس نام محصول مورد نظر خود را وارد کنید تا جستجو صورت پذیرد
		`), chatID)
		return
	}
	searchQ := strings.Join(splitMsg[1:], " ")

	var products []models.Product
	err := database.
		DB().
		Select("title").
		From("products").
		Where(dbx.Like("title", searchQ)).
		OrWhere(dbx.NewExp("title = {:title}", dbx.Params{"title": searchQ})).
		All(&products)
	if err != nil {
		handleErr(err, chatID)
		return
	}

	if len(products) == 0 {
		handleBotMessage("پس از جستجو موردی جهت نمایش یافت نشد", chatID)
		return
	}

	productsKeyboard(products, chatID)

}

func helpCommand(chatID int64) {
	text := `
	راهنما:

	با استفاده از ربات فروشگاه ما میتوانید همیشه به جدیدترین محصولات ما دسترسی داشته باشید و اطلاعاتی شامل قیمت، عکس هایی با کیفیت های مختلف، ویدئوی محصول و دیگر اطلاعات خاص محصولات را مشاهده نمایید!

	توضیحات مربوط به دستورات ربات:

	۱. /start
	با این دستور ربات برای شما ریستارت میشود

	۲. /newest
	این دستور جدیدترین محصولات ثبت شده در فروشگاه را به شما نشان میدهد

	۳. /search
	با این دستور میتوانید محصول مورد نظر خود را جستجو نمایید، برای توضیحات بیشتر، این دستور را به ربات ارسال کنید

	۴. /help
	توضیحات فعلی را با این دستور میتوانید مشاهده نمایید
	`

	handleBotMessage(text, chatID)
}

func defaultCommand(chatID int64) {
	handleBotMessage("دستور وارد شده نامعتبر است", chatID)
}
