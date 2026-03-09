package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	// Webサーバーを起動
	log.Fatal(http.ListenAndServe(":5000", server))
}