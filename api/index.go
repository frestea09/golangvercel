// api/index.go
package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}


func main() {

    http.HandleFunc("/", Handler)

    fmt.Println("Server is running on http://localhost:3000")

    if err := http.ListenAndServe(":3000", nil); err != nil {

        fmt.Println(err)

    }

}