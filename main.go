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

func checkCORS(w http.ResponseWriter, r *http.Request) (ok bool) {
  ok = false
  if r.Method == http.MethodOptions {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Max-Age", "3600")
    w.WriteHeader(http.StatusNoContent)
    ok = true
  }
  return
}
