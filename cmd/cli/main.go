package main

import (
	"fmt"
	poker "github.com/takaya-47/go_by_tdd_application"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeFunc()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewGame(poker.BlindAlerterFunc(poker.Alerter), store)
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
