package poker_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	poker "github.com/takaya-47/go_by_tdd_application"
)

func TestGetPlayers(t *testing.T) {
	store := poker.NewStubPlayerStore(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}, nil)

	game := &GameSpy{}
	server := mustMakePlayerServer(t, store, game)

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
	game := &GameSpy{}
	server := mustMakePlayerServer(t, store, game)

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
		game := &GameSpy{}
		server := mustMakePlayerServer(t, store, game)

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
		game := &GameSpy{}
		server := mustMakePlayerServer(t, store, game)

		request := poker.NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response, http.StatusOK)
	})

	t.Run("start a game with 3 players, send some blind alerts down WS and declare Ruth the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "Ruth"

		store := poker.NewStubPlayerStore(nil, nil)
		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := mustDialWS(t, wsURL)
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, winner)

		within(t, 10*time.Millisecond, func() { assertWebsocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game *GameSpy) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func assertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf(`got "%s", want "%s"`, string(msg), want)
	}
}

func within(t *testing.T, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	// assert関数とtime.Afterの間でselectして、assert関数が完了するか、タイムアウトするかを待つ
	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}
