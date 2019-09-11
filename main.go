package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// TODO: add config
// TODO: build a workpool
func main() {
	http.HandleFunc("/", SignImage)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
