package poker_test

import (
	poker "github.com/takaya-47/go_by_tdd_application"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := poker.CreateTempFile(t, "[]")
	defer cleanDatabase()

	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)

	game := &GameSpy{}
	server, _ := poker.NewPlayerServer(store, game)
	player := "Pepper"

	for i := 0; i < 3; i++ {
		server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(player))

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewLeagueRequest())

		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
