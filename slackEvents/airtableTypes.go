package slackevents

// ConfirmedPeople is a struct of the airtable record format
type ConfirmedPeople struct {
	AirtableID string
	Fields     struct {
		Name     string
		Email    string
		SlackID  string
		Mentor   []string
		mentorID string
	}
}
