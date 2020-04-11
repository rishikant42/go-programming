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

	mf := &MediaFile{
		Link: "www.Link.com",
	}
	tp := TemplateParameter{
		Type:     "header",
		Document: mf,
	}

	tc := TemplateComponent{
		Type: "TemplateComponent",
	}
	tc.Parameters = append(tc.Parameters, &tp)

	templateMsg := TemplateMessage{}
	templateMsg.Template.Components = append(templateMsg.Template.Components, &tc)
	hc, _ := json.Marshal(templateMsg)
	// mfg, _ := json.Marshal(mediaFileMsg)
	// mbm, _ := json.Marshal(messageBaseMsg)
	fmt.Println(string(hc))
}
