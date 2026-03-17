package poker_test

import (
	poker "github.com/takaya-47/go_by_tdd_application"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("works with empty file", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)

		poker.AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{"Chris", 33}, // スコアが高い順
			{"Cleo", 10},
		}

		poker.AssertLeague(t, got, want)

		// 2回目以降も同じ結果が得られることを確認
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33

		poker.AssertScoreEquals(t, got, want)
		poker.AssertNoError(t, err)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		poker.AssertScoreEquals(t, got, want)
		poker.AssertNoError(t, err)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		poker.AssertScoreEquals(t, got, want)
		poker.AssertNoError(t, err)
	})
}
