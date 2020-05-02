package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/brianloveswords/airtable"
	"github.com/joho/godotenv"
	"github.com/neelr/meeting-coordinator/slackevents"
)

func main() {
	fmt.Println("Starting up the bot....")
	godotenv.Load()

	// Create and link airtable base
	slackevents.Base = airtable.Client{
		APIKey: os.Getenv("AIRTABLE_KEY"),
		BaseID: os.Getenv("BASE"),
	}
	// API Routes
	http.HandleFunc("/api/createMeeting", slackevents.CreateMeeting)
	http.HandleFunc("/api/blocks", slackevents.Blocks)

	fmt.Println("Up on port 3000!")
	http.ListenAndServe(":3000", nil)
}
