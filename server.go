package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
)

func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi")
  })

  http.HandleFunc("/beau", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><body><h1>Beau</h1></body></html>")
  })

  log.Fatal(http.ListenAndServe(":8081", nil))

}
