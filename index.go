package main

import (
	"fmt"
	"net/http"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"hello world")
	})
	if err:= http.ListenAndServe(":8000", mux); err != nil{
		fmt.Print(err.Error())
	}
}