package main

import (
	"encoding/json"
	"fmt"
)

type waJobBase struct {
	To          string `json:"to,omitempty"`
	From        string `json:"from,omitempty"`
	MessageType string `json:"type,omitempty"`
	Uid         string `json:"uid"`
}

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

type TemplateParameter struct {
	Type     string     `json:"type,omitempty"`
	Text     string     `json:"text,omitempty"`
	Document *MediaFile `json:"document,omitempty"`
	Image    *MediaFile `json:"image,omitempty"`
	Video    *MediaFile `json:"video,omitempty"`

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
}

type TemplateComponent struct {
	Type string `json:"type,omitempty"`

	Parameters []*TemplateParameter `json:"parameters,omitempty"`
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

		Name string `json:"name,omitempty"`

		Components []*TemplateComponent `json:"components,omitempty"`
	} `json:"template,omitempty"`
}

func main() {
	msg := []byte(`

{
    "to": "your-test-recipient-wa-id",
    "recipient_type": "individual",
    "type": "template",
    "template": {
	        "namespace": "88b39973_f0d5_54e1_29cf_e80f1e3da4f2",
        "name": "flight_confirmation",
        "language": {
            "code": "en",
            "policy": "deterministic"
        },
        "components": [
            {
                "type": "header",
                "parameters": [
                    {
                        "type": "document",
                        "document": {
                        	"filename": "MRRATH-CGK-KUL.pdf",
                        	"link": "link-to-your-document"
                        }
                    }
                ]
            },
            {
                "type": "body",
                "parameters": [
                	{
                		"type": "text",
                		"text": "CGK (Jakarta)"
                	},
                	{
                		"type": "text",
                		"text": "KUL (Kuala Lumpur)"
                	},
                    {
                    	"type": "date_time",
                        "date_time" : {
                            "fallback_value": "20th April 2019, 12:20pm",
                            "day_of_month": "20",
                            "year": "2019",
                            "month": "9",
                            "hour": "12",
                            "minute": "10"
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
	// mfg, _ := json.Marshal(mediaFileMsg)
	// mbm, _ := json.Marshal(messageBaseMsg)
	fmt.Println(string(hc))
}
