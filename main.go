package main

import "net/http"

func main() {
	// PlayerServer関数をhttp.HandlerFunc型にキャスト
	handler := http.HandlerFunc(PlayerServer)
	http.ListenAndServe(":5000", handler)
}