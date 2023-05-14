package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		message := Message{Text: "Hello, World!"}
		json.NewEncoder(w).Encode(message)
	})
	http.ListenAndServe(":8080", nil)
}
