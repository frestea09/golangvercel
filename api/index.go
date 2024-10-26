// api/index.go

package handler

import (
	"fmt"
	"net/http"
)

// Handler adalah fungsi yang menangani permintaan HTTP

func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

}
