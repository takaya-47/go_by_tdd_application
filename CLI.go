package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func (cli *CLI) PlayPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.store.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}