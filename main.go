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
