package slackapi

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// SendMessage sends a Slack Message through the WebAPI with a text, channle, and an optional block
func SendMessage(text string, channel string, blocks interface{}) string {
	data := url.Values{}
	data.Set("channel", channel)
	data.Set("token", os.Getenv("SLACK_API_TOKEN"))
	switch blocks.(type) {
	case string:
		data.Set("blocks", blocks.(string))
	case nil:
		data.Set("text", text)
	}
	client := &http.Client{}
	r, _ := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(r)
	byteBody, _ := ioutil.ReadAll(resp.Body)
	return string(byteBody)
}
