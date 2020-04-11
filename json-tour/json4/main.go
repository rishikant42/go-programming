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
	// fmt.Println("Hello world")
	payload := map[string]interface{}{
		"queue_type":    "wa",
		"from":          "FROM",
		"to":            "TO",
		"uid":           "UID",
		"account_token": "AccountTOKEN",
		"account_uid":   "accounYUD",
	}
	payload["type"] = "media_template"

	payload["attachments"] = []string{"image", "video"}
	payload["media_url"] = "www.media-url.com"

	templateParams := []string{"param1", "param2"}

	body_params := []map[string]string{}

	for _, param := range templateParams {
		body_params = append(body_params, map[string]string{
			"type": "text",
			"text": param,
		})
	}

	components := []map[string]interface{}{
		{
			"type":       "body",
			"parameters": body_params,
		},
	}

	payload["template"] = map[string]interface{}{
		"namespace": "c.templated_text.Namespace",
		"language": map[string]interface{}{
			"policy": "deterministic",
			"code":   "c.templated_text.LanguageCode",
		},
		"name":       "c.templated_text.Name",
		"components": components,
	}

	hc, _ := json.Marshal(payload)
	fmt.Println(string(hc))

}
