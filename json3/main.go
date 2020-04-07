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

		Name string `json:"name,omitempty"`

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
   "account_uid" : "accounYUD",
   "template" : {
      "namespace" : "c.templated_text.Namespace",
      "name" : "c.templated_text.Name",
      "language" : {
         "code" : "c.templated_text.LanguageCode",
         "policy" : "deterministic"
      },
      "components" : [
         {
            "parameters" : [
               {
                  "type" : "text",
                  "text" : "param1"
               },
               {
                  "type" : "text",
                  "text" : "param2"
               }
            ],
            "type" : "body"
         }
      ]
   },
   "uid" : "UID",
   "to" : "TO",
   "from" : "FROM",
   "account_token" : "AccountTOKEN",
   "attachments" : [
      "image",
      "video"
   ],
   "type" : "template",
   "queue_type" : "wa"
 }

	`)

	templateMsg := TemplateMessage{}

	json.Unmarshal(msg, &templateMsg)
	hc, _ := json.Marshal(templateMsg)
	fmt.Println(string(hc))
}
