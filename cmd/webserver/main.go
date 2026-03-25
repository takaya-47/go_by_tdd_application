package main

import (
	poker "github.com/takaya-47/go_by_tdd_application"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeFunc()

	server, err := poker.NewPlayerServer(store)

	if err != nil {
		log.Fatalf("could not create player server %v", err)
	}

	// Webサーバーを起動
	if err := http.ListenAndServe(":5001", server); err != nil {
		log.Fatalf("could not listen on port 5001 %v", err)
	}
}