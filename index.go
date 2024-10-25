package main

import (
	"fmt"
	"net/http"
)

// Handler untuk endpoint "/comment"
func commentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/comment", commentHandler)

	// Vercel akan menjalankan fungsi ini, jadi tidak perlu menjalankan server secara manual
}