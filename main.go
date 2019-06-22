package serverless

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

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
	name := r.FormValue("name")
	if name == "" {
		name = "next"
	}
	var m image.Image
	m, err = jpeg.Decode(imageBuffer)
	if err != nil {
		panic(err)
	}
	watermark := signature(name)
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
