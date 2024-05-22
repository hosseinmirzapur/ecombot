package main

import (
	"log"

	"github.com/hosseinmirzapur/ecombot/bot"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	godotenv.Load()

	// load telegram bot
	err := bot.RegisterBot()
	if err != nil {
		log.Fatal(err)
	}
}
