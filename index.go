package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "text/html") // Set the correct content type

    fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

}