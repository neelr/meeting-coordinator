package slackevents

import "github.com/brianloveswords/airtable"

// ConfirmedPeople is a struct of the airtable record format
type ConfirmedPeople struct {
	airtable.Record
	Fields struct {
		Name     string   `json:",omitempty"`
		Email    string   `json:",omitempty"`
		SlackID  string   `json:",omitempty"`
		Mentor   []string `json:",omitempty"`
		MentorID []string `json:",omitempty"`
	}
}

// Meetings for the airtable record
type Meetings struct {
	airtable.Record
	Fields struct {
		MeetingID     string   `json:",omitempty"`
		MeetingDate   string   `json:",omitempty"`
		Mentee        []string `json:",omitempty"`
		Status        string   `json:",omitempty"`
		MentorSlackID []string `json:",omitempty"`
		MenteeSlackID []string `json:",omitempty"`
		Note          string   `json:",omitempty"`
	}
}
