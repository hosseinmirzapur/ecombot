package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/hosseinmirzapur/ecombot/bot"
	"github.com/hosseinmirzapur/ecombot/database"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	godotenv.Load()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(1)
	}()

	// load DB
	go pocketbaseDB()

	// run telegram bot
	tgRun()

}

func tgRun() {
	//load telegram bot
	err := bot.RegisterBot()
	if err != nil {
		log.Fatalln(err)
	}

	// register bot commands
	err = bot.RegisterCommands()
	if err != nil {
		log.Fatalln(err)
	}

	// set bot debug mode
	bot.SetDebug(true)

	// listen for updates
	bot.ListenForUpdates()
}

func pocketbaseDB() {
	err := database.RegisterDB()
	if err != nil {
		log.Fatal(err)
	}
}
