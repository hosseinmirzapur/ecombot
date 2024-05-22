package main

import (
	"log"

	"github.com/hosseinmirzapur/ecombot/bot"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	godotenv.Load()

	// load DB
	err := database.RegisterDB()
	if err != nil {
		log.Fatal(err)
	}

	// auto-migrate
	err = database.AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}

	// load telegram bot
	err = bot.RegisterBot()
	if err != nil {
		log.Fatalln(err)
	}

	// set bot debug mode
	bot.SetDebug(true)
}
