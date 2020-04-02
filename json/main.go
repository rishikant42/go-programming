package main

import (
	"encoding/json"
	"fmt"
)

type MediaFile struct {
	Link     string `json:"link,omitempty"`
	Provider struct {
		Name string `json:"name,omitempty"`
	} `json:"provider,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type HeaderComponent struct {
	Type       string `json:"type,omitempty"`
	Parameters []struct {
		Type string `json:"type,omitempty"`

		Text string `json:"text,omitempty"`

		Document *MediaFile `json:"document,omitempty"`

		Image *MediaFile `json:"image,omitempty"`
	} `json:"parameters,omitempty"`
}

func main() {
	media := MediaFile{
		Link: "www.google.com",
		// Caption: "World",
	}
	media.Provider.Name = "Hello"

	mediaData, err := json.Marshal(media)

	if err != nil {
		fmt.Println("Error occured:", err)
	}

	fmt.Println(string(mediaData))

	headerData := []byte(`
            {
                "type": "header",
								"parameters": [
									{
										"type": "text",
										"document": {
											"link": "www.google.com"
										}
									}
								]
            }
	`)

	var headerComponent HeaderComponent

	errNew := json.Unmarshal(headerData, &headerComponent)

	if errNew != nil {
		fmt.Println("Error occured:", errNew)
	}

	// fmt.Printf("%+v\n", headerComponent)

	hc, _ := json.Marshal(headerComponent)
	fmt.Println(string(hc))
}
