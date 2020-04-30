package slackevents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/neelr/meeting-coordinator/slackapi"
	"github.com/fabioberger/airtable-go"
)

// MainHandle handles all requests from the server, and redirects them to the correct place
func MainHandle(w http.ResponseWriter, r *http.Request) {
	client, err := airtable.New(os.Getenv("AIRTABLE_KEY") , os.Getenv("BASE"))
	if err != nil {
		panic(err)
	}
	records := []ConfirmedPeople{}
	if err := client.ListRecords("Confirmed People", &records); err != nil {
		panic(err)
	}
	fmt.Println(records[0])
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var jsonBody map[string]string
	json.Unmarshal(b, &jsonBody)
	slackapi.SendMessage("wow this is me and not a bot ahahahahahhaha", "C0P5NE354", nil)
	fmt.Fprintf(w, jsonBody["challenge"])
}

func slashSetMeeting() {

}
