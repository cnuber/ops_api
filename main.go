package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q",
html.EscapeString(r.URL.Path))
    })
    log.Fatal(http.ListenAndServe(":8080", nil))

}
