package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/neelr/meeting-coordinator/slackevents"
)

func main() {
	fmt.Println("Starting up the bot....")
	godotenv.Load()
	http.HandleFunc("/api/events", slackevents.MainHandle)

	fmt.Println("Up on port 3000!")
	http.ListenAndServe(":3000", nil)
}
