package slack

import (
	"time"
	"errors"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
)

type Auth struct {
	Token string `json:"token"`
}

type Parameters struct {
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments"`
	Username    string       `json:"username"`
	Iconemoji   string       `json:"icon_emoji"`
}

type Attachment struct {
	Color  string  `json:"color"`
	Title  string  `json:"title"`
	Text   string  `json:"text"`
	Fields []Field `json:"fields"`
	Footer string  `json:"footer"`
	Ts     int64   `json:"ts"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

const PostMessageUrl = "https://slack.com/api/chat.postMessage"
const AuthFile = "slack.json"

var slackAuth Auth

func Init(filePath string) (error) {
	correctPath := fixPath(filePath)
	fileContent, err := ioutil.ReadFile(correctPath + AuthFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContent, &slackAuth)
	if err != nil {
		return err
	}

	return nil
}

func fixPath(path string) string {
	lastChar := path[len(path)-1:]
	if lastChar != "/" {
		path += "/"
	}
	return path
}

func setAttDefaults(a *Attachment) {
	a.Color = "#ff000"
	a.Text = ""
	a.Footer = ""
	a.Fields = []Field{}
	a.Ts = time.Now().Unix()
}

func setParamDefaults(p *Parameters) {
	p.Username = "Web Team"
	p.Iconemoji = ":joy:"
}

func setParameters(username string, iconEmoji string, channel string, attachment Attachment) (Parameters, error) {
	p := Parameters{}
	setParamDefaults(&p)

	if channel == "" {
		return Parameters{}, errors.New("slack: channel attribute can not be empty")
	} else {
		if !strings.Contains(channel, "#") {
			channel = "#" + channel
		}
	}

	if username != "" {
		p.Username = username
	}
	if iconEmoji != "" {
		p.Iconemoji = iconEmoji
	}

	p.Channel = channel
	attachments := []Attachment{attachment}
	p.Attachments = attachments

	return p, nil
}

func MakeAttachment(color string, title string, text string, fields []Field, timestamp int64) (Attachment, error) {
	a := Attachment{}
	setAttDefaults(&a)

	if color != "" {
		a.Color = color
	}

	if title != "" {
		a.Title = title
	}

	if text != "" {
		a.Text = text
	}

	if len(fields) == 0 {
		return Attachment{}, errors.New("slack: fields array attribute can not be empty")
	}

	if timestamp != 0 {
		a.Ts = timestamp
	}

	a.Fields = fields

	return a, nil
}

func MakeField(title string, value string) (Field, error) {
	f := Field{}
	f.Title = title
	f.Value = value

	return f, nil
}

func SendMessage(username string, iconEmoji string, channel string, attachment Attachment) (bool, error) {
	parameters, err := setParameters(username, iconEmoji, channel, attachment)
	if err != nil {
		return false, err
	}

	jsonBodyParameters, err := json.Marshal(parameters)
	if err != nil {
		return false, err
		// log.Fatal(err)
	}

	url := PostMessageUrl
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBodyParameters))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", "Bearer "+slackAuth.Token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	result, err := parseBodyResponse(body)
	if err != nil {
		return false, err
	}

	return result, nil
}

func parseBodyResponse(body []byte) (bool, error) {
	result := fmt.Sprintf("%s", body)

	switch result {
	case "no_text":
		return false, errors.New("slack: no text error")
	}
	return true, nil
}
