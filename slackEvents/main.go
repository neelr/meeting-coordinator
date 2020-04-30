package slackevents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/neelr/meeting-coordinator/slackapi"
)

// MainHandle handles all requests from the server, and redirects them to the correct place
func MainHandle(w http.ResponseWriter, r *http.Request) {
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
