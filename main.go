package main

import "net/http"

func main() {
	handler := http.HandlerFunc(PlayerServer)
	http.ListenAndServe(":5000", handler)
}