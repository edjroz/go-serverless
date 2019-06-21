package serverless

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"
  "imageUri"
  "io"
)

var (
  imageUri string
  image io.ReadCloser
)
func init() {
  if baseUri == nil {
    imageUri = os.GetEnv("BASE_IMAGE_URI")
  }
  if image == nil {
    var err error
    imagePayload, err := http.Get(baseUri)
    if err != nil {
      fmt.Fprintf(w, "Error %d", err)
      return
    }
    defer imagePayload.Body.Close()
    image = imagePayload.Body
  }
}
func SignImage(w http.ResponseWriter, r *http.Request) {
  // TODO: how should your data be??
  cors := checkCORS(w, r); if cors {
    return
  }
}
