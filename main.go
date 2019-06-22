package serverless

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

var defaultWatermark image.Image

func init () {
	if defaultWatermark == nil {
		defaultWatermark = signature("Google")
	}
}
func SignImage(w http.ResponseWriter, r *http.Request) {
	cors := checkCORS(w, r)
	if cors {
		return
	}
	imageBuffer, _, err := r.FormFile("picture")
	if err != nil {
		panic(err)
	}
	defer imageBuffer.Close()

	var m image.Image
	m, err = jpeg.Decode(imageBuffer)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	var watermark image.Image
	switch name {
	  case "":
			watermark = defaultWatermark
		default:
			watermark = signature(name)
	}

	signed := drawImage(m, watermark)

	var buff bytes.Buffer
	err = jpeg.Encode(&buff, signed, nil)
	if err != nil {
		panic(err)
	}
	if _, err = io.Copy(w, &buff); err != nil {
		panic(err)
	}
}
