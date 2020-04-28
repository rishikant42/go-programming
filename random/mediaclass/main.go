package main

import (
	"errors"
	"fmt"
	"mime"
	"net/http"
	"strings"
)

const (
	Doc   = "document"
	Image = "image"
	Video = "video"
)

var ContentTypeMediaClass = map[string]string{
	"application/pdf": Doc,
	"image/jpeg":      Image,
	"image/jpg":       Image,
	"image/png":       Image,
	"video/mp4":       Video,
}

func getMediaClass(url string) (string, error) {

	mresp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer mresp.Body.Close()

	if len(mresp.Header["Content-Type"]) == 0 {
		return "", errors.New("whatsapp.ErrMediaContentTypeNotSet")
	}

	contentType := strings.ToLower(mresp.Header["Content-Type"][0])

	mType, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		return "", errors.New("whatsapp.ErrMediaNotSupported")
	}
	mediaClass, ok := ContentTypeMediaClass[mType]

	if !ok {
		fmt.Println("I am here")
		return "", errors.New("whatsapp.ErrMediaNotSupported")
	}

	return mediaClass, nil

}

func main() {
	url := "https://storage.googleapis.com/wa_media/whatsappvideo.3gp"
	mClass, err := getMediaClass(url)
	fmt.Println("url:", url)
	fmt.Println("mediaClass:", mClass)
	fmt.Println("Error:", err)
}
