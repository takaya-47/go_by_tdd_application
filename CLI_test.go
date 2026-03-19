package poker_test

import (
	"bytes"
	"fmt"
	poker "github.com/takaya-47/go_by_tdd_application"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players, starts the game and finishes the game", func(t *testing.T) {
		in := strings.NewReader("7\nChris wins\n")
		stdOut := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		got := stdOut.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}

		if game.FinishedWith != "Chris" {
			t.Errorf("wanted Finish called with 'Chris' but got %q", game.FinishedWith)
		}

	})
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

type GameSpy struct {
	StartedWith int
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}