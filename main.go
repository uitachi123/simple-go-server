package main

import (
	"io"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to new server!")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Goodbye")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/goodbye", goodbye)
	// listen to port
	http.ListenAndServe(":8080", mux)
}
