package serverless

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
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
