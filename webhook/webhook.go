package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

type Embed struct {
	Color     int       `json:"color"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Author    Author    `json:"author"`
	Title     string    `json:"title"`
	Timestamp string    `json:"timestamp"`
	Fields    []Fields  `json:"fields"`
	Footer    Footer    `json:"footer"`
}

type Thumbnail struct {
	Url string `json:"url"`
}

type Footer struct {
	Text     string `json:"text"`
	Icon_url string `json:"icon_url"`
}

type Author struct {
	Name     string `json:"name"`
	Icon_url string `json:"icon_url"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

func Create(imageUrl string, title string, date string, version string) Webhook {
	color := 16318296
	webhook := Webhook{
		Content: "",
		Embeds: []Embed{
			Embed{
				Color: color,
				Thumbnail: Thumbnail{
					Url: imageUrl,
				},
				Title:     title,
				Timestamp: date,
				Fields:    []Fields{},
				Footer: Footer{
					Text:     "Volt Scripts - v. " + version,
					Icon_url: "https://i.postimg.cc/vB3MDK2s/t-pfp.png",
				},
			},
		},
	}
	return webhook
}

func (webhook Webhook) AddField(name string, value string, inline bool) {
	newField := Fields{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
	webhook.Embeds[0].Fields = (append(webhook.Embeds[0].Fields, newField))
}

func (webhook Webhook) Send(url string) (*http.Response, error) {
	client := &http.Client{}

	webhookData, err := json.Marshal(webhook)
	if err != nil {
		fmt.Println(err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(webhookData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err.Error())
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	return res, err
}
