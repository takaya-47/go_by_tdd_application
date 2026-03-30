package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	poker "github.com/takaya-47/go_by_tdd_application"
)

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		in := userSends(t, "3", "Chris wins")
		stdOut := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		in := userSends(t, "8", "Cleo wins")
		stdOut := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 8)
		assertGameFinishedWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := userSends(t, "Pies")
		stdOut := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdOut, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		in := userSends(t, "8", "Lloyd is a killer")
		stdOut := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, stdOut, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})
}

func userSends(t *testing.T, messages ...string) *strings.Reader {
	t.Helper()

	message := strings.Join(messages, "\n")
	return strings.NewReader(message)
}

func assertGameStartedWith(t *testing.T, game *GameSpy, want int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == want
	})

	if !passed {
		t.Errorf("wanted Start called with %d but got %d", want, game.StartedWith)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameFinishedWith(t *testing.T, game *GameSpy, want string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == want
	})

	if !passed {
		t.Errorf("expected finish called with %q but got %q", want, game.FinishedWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)

	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()

	if game.FinishCalled {
		t.Error("game should not have finished")
	}
}

func assertMessagesSentToUser(t *testing.T, stdOut *bytes.Buffer, messages ...string) {
	t.Helper()

	got := stdOut.String()
	want := strings.Join(messages, "")

	if got != want {
		t.Errorf("got %v sent to stdout but expected %+v", got, messages)
	}
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

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	BlindAlert   []byte
	FinishCalled bool
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}
