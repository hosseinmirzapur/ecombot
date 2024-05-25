package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/hosseinmirzapur/ecombot/bot"
	"github.com/hosseinmirzapur/ecombot/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(1)
	}()

	_, err := w.Write([]byte("deployed successfully!"))
	if err != nil {
		fmt.Fprintf(w, "<h1>Failed to write to output</h1>")
	}

	// load DB
	go PocketbaseDB()

	// run telegram bot
	TgRun()

}

func TgRun() {
	//load telegram bot
	err := bot.RegisterBot()
	if err != nil {
		log.Println(err)
		return
	}

	// register bot commands
	err = bot.RegisterCommands()
	if err != nil {
		log.Println(err)
		return
	}

	// set bot debug mode
	bot.SetDebug(true)

	// listen for updates
	bot.ListenForUpdates()
}

func PocketbaseDB() {
	err := database.RegisterDB()
	if err != nil {
		log.Fatal(err)
	}
}
