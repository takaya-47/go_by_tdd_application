package poker_test

import (
	poker "github.com/takaya-47/go_by_tdd_application"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := poker.NewStubPlayerStore(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}, nil)

	server := poker.NewPlayerServer(store)

	t.Run("return Peppers's score", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := poker.NewStubPlayerStore(nil, nil)
	server := poker.NewPlayerServer(store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := poker.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusAccepted)
		poker.AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := poker.League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.NewStubPlayerStore(nil, wantedLeague)
		server := poker.NewPlayerServer(store)

		request := poker.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := poker.GetLeagueFromResponse(t, response.Body)

		poker.AssertStatus(t, response, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response, "application/json")
	})
}

func TestGame(t *testing.T) {
	t.Run("Get /game returns 200", func(t *testing.T) {
		store := poker.NewStubPlayerStore(nil, nil)
		server := poker.NewPlayerServer(store)

		request := poker.NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})
}
