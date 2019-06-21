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
  var data struct {
    Message string `json:"message"`
  }

  cors := checkCORS(w, r); if cors {
    return
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")

  if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
    fmt.Fprint(w, "Hello World!")
    return
  }
  if data.Message == "" {
    fmt.Fprint(w, "Hello World!")
    return
  }
  fmt.Fprint(w, html.EscapeString(data.Message))
}
