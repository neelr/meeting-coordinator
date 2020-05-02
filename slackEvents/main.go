package slackevents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/brianloveswords/airtable"
	"github.com/neelr/meeting-coordinator/slackapi"
	"github.com/tj/go-naturaldate"
)

// Base is the airtable base
var Base airtable.Client

const (
	layoutUS = "January 2, 2006"
)

// CreateMeeting handles the slash command from the slack
func CreateMeeting(w http.ResponseWriter, r *http.Request) {
	confirmedPeopleTable := Base.Table("Confirmed People")
	meetingsTable := Base.Table("Meetings")
	fmt.Println("logs> /create-meeting command")

	// Parse the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	query, _ := url.ParseQuery(string(b))

	// Search the table for their slackID
	records := []ConfirmedPeople{}
	confirmedPeopleTable.List(&records, &airtable.Options{
		Filter: fmt.Sprintf(`{SlackID} = "%s"`, query["user_id"][0]),
	})

	// If returns no results, then not authorized
	if len(records) == 0 {
		fmt.Fprintf(w, "Are you a signed /up as a Person for Hack Club Summer of Making? If so DM <@UJYDFQ2QL>")
		return
	}

	// Create Meeting Record
	queryArray := strings.Split(query["text"][0], "|")
	t, err := naturaldate.Parse(queryArray[0], time.Now())
	newMeeting := &Meetings{}
	newMeeting.Fields.Mentee = []string{records[0].ID}
	newMeeting.Fields.MeetingDate = t.Format("2006-01-02T15:04:05.999Z") // Format to ISO 8601
	newMeeting.Fields.Status = "Pending"
	newMeeting.Fields.Note = queryArray[1]

	// Check if parsed time is correct
	sunday, _ := naturaldate.Parse("Next Sunday at 11:59PM", time.Now())
	if t.After(sunday) || t.Before(time.Now()) || err != nil {
		fmt.Println(t)
		fmt.Fprintf(w, fmt.Sprintf("Time error! Make sure that its after today, and  before this Sunday! The date parsed was `%s`", t))
		return
	}
	err = meetingsTable.Create(newMeeting)
	fmt.Println(err)

	// Send Message with Block
	fmt.Fprintf(w, fmt.Sprintf("Creating Meeting, and Sent to Mentor with the note `%s`!", queryArray[1]))
	slackapi.SendMessage(
		"",
		records[0].Fields.MentorID[0],
		fmt.Sprintf(
			`[
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "Hi! <@%s> has scheduled a meeting with you at %s for the reason '%s'"
				},
				"accessory": {
					"type": "button",
					"text": {
						"type": "plain_text",
						"text": "Accept Meeting!",
						"emoji": true
					},
					"style": "primary",
					"value": "Accepted|%s|%s|%s"
				}
			}
		]`, // In Block, the values are the recordID, Mentee slackID, and the date
			records[0].Fields.SlackID,
			t.Format("Monday 3:04 PM"),
			queryArray[1],
			newMeeting.ID,
			records[0].Fields.SlackID,
			t.Format("Monday 3:04 PM"),
		),
	)
}

// Blocks handles slack block evvents
func Blocks(w http.ResponseWriter, r *http.Request) {
	meetingsTable := Base.Table("Meetings")
	fmt.Println("logs> block request")

	// Parse the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Look at payload and unwrap data
	query, _ := url.ParseQuery(string(b))
	var payload map[string]interface{}
	json.Unmarshal([]byte(query["payload"][0]), &payload)
	value := payload["actions"].([]interface{})[0].(map[string]interface{})["value"].(string)
	id := payload["user"].(map[string]interface{})["id"].(string)

	// Get params from value
	params := strings.Split(value, "|")

	// Create and update record to whatever the Block said to
	record := Meetings{}
	record.ID = params[1]
	record.Fields.Status = params[0]
	meetingsTable.Update(&record)

	// Send slack messages to Mentee and Mentor
	slackapi.SendMessage(fmt.Sprintf("Your meeting was accepted! Be ready at %s.", params[3]), params[2], nil)
	slackapi.SendMessage(fmt.Sprintf("You accepted the meeting! Be ready at %s.", params[3]), id, nil)
}
