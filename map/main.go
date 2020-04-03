package main

import "fmt"

const (
	MaxMediaLength = 64000000
	Doc            = "document"
	Image          = "image"
	Audio          = "audio"
	Video          = "video"
)

var ValidMimeSupported = map[string]string{
	"application/pdf":    Doc,
	"application/msword": Doc,
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   Doc,
	"application/vnd.ms-powerpoint":                                             Doc,
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": Doc,
	"application/vnd.ms-excel":                                                  Doc,
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         Doc,
	"image/jpeg": Image,
	"image/jpg":  Image,
	"image/png":  Image,
	"audio/amr":  Audio,
	"audio/mpeg": Audio,
	"audio/aac":  Audio,
	"audio/m4a":  Audio,
	"audio/ogg":  Audio,
	"video/mp4":  Video,
	"video/3gpp": Video,
}

func tempMain() (string, error) {
	mType := "application/msword"
	mediaClass, ok := ValidMimeSupported[mType]
	if !ok {
		fmt.Println("Hello", mediaClass)
		fmt.Println("ok", ok)
	} else {
		fmt.Println("world", mediaClass)
		fmt.Println("ok", ok)
	}
	return mediaClass, nil
}

func main() {
	value, err := tempMain()
	fmt.Println("value: ", value)
	fmt.Println("err: ", err)
}
