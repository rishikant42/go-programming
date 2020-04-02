package main

import (
	"encoding/json"
	"fmt"
)

type MessageBase struct {
	RecipientType string `json:"recipient_type,omitempty"`
	To            string `json:"to,omitempty"`
	MessageType   string `json:"type,omitempty"`
}

type MediaFile struct {
	Link string `json:"link,omitempty"`

	Provider struct {
		Name string `json:"name,omitempty"`
	} `json:"provider,omitempty"`

	FileName string `json:"filename,omitempty"`
}

type TemplateMessage struct {
	MessageBase
	TTL string `json:"ttl,omitempty"`

	Template struct {
		Namespace string `json:"namespace,omitempty"`

		Language struct {
			Policy string `json:"policy,omitempty"`
			Code   string `json:"code,omitempty"`
		} `json:"language,omitempty"`

		Name string `json:"names,omitempty"`

		Components []struct {
			Type string `json:"type,omitempty"`

			Parameters []struct {
				Type     string     `json:"type,omitempty"`
				Text     string     `json:"text,omitempty"`
				Document *MediaFile `json:"document,omitempty"`
				Image    *MediaFile `json:"image,omitempty"`

				Currency *struct {
					FallbackValue string `json:"fallback_value,omitempty"`
					Code          string `json:"code,omitempty"`
					Amount        string `json:"amount_1000,omitempty"`
				} `json:"currency,omitempty"`

				DateTime *struct {
					FallbackValue string `json:"fallback_value,omitempty"`
					DayOfWeek     string `json:"day_of_week,omitempty"`
					DayOfMonth    string `json:"day_of_month,omitempty"`
					Year          string `json:"year,omitempty"`
					Month         string `json:"month,omitempty"`
					Hour          string `json:"hour,omitempty"`
					Minute        string `json:"minute,omitempty"`
					Timestamp     string `json:"timestamp,omitempty"`
				} `json:"date_time,omitempty"`
			} `json:"parameters,omitempty"`
		} `json:"components,omitempty"`
	} `json:"template,omitempty"`
}

func main() {
	msg := []byte(`

{
	"to": "recipient_wa_id",
	"ttl": "P4D",
	"type": "template",
	"template": {
		"namespace": "your-namespace",
		"language": {
			"policy": "deterministic",
			"code": "your-language-and-locale-code"
		},
		"name": "your-template-name",
		"components": [{
				"type": "header",
				"parameters": [{
						"type": "header",
						"text": "replacement-text"
					},
					{
						"type": "document",
						"document": {
							"link": "the-provider-name/protocol://the-url",
							"provider": {
								"name": "provider-name"
							},
							"filename": "your-document-filename"
						}
					},
					{
						"type": "image",
						"document": {
							"link": "the-provider-name/protocol://the-url",
							"provider": {
								"name": "provider-name"
							},
							"filename": "your-document-filename"
						}
					}
				]
			},
			{
				"type": "body",
				"parameters": [{
						"type": "text",
						"text": "replacement_text"
					},
					{
						"type": "currency",
						"currency": {
							"fallback_value": "$100.99",
							"code": "USD",
							"amount_1000": "100990"
						}
					},
					{
						"type": "date_time",
						"date_time": {
							"fallback_value": "February 25, 1977",
							"day_of_week": "5",
							"day_of_month": "25",
							"year": "1977",
							"month": "2",
							"hour": "15",
							"minute": "33",
							"timestamp": "1485470276"
						}
					}
				]
			}
		]
	}
}
	`)

	templateMsg := TemplateMessage{}
	mediaFileMsg := MediaFile{}
	messageBaseMsg := MessageBase{}

	errNew := json.Unmarshal(msg, &templateMsg)
	errNew = json.Unmarshal(msg, &mediaFileMsg)
	errNew = json.Unmarshal(msg, &messageBaseMsg)

	if errNew != nil {
		fmt.Println("Error occured:", errNew)
	}

	hc, _ := json.Marshal(templateMsg)
	mfg, _ := json.Marshal(mediaFileMsg)
	mbm, _ := json.Marshal(messageBaseMsg)
	fmt.Println(string(hc))
	fmt.Printf("\n\n")
	fmt.Println("Hello world")
	fmt.Println(string(mfg))
	fmt.Println(string(mbm))
}
