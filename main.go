package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
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

func main() {
	http.HandleFunc("/", SignImage)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
