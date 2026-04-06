package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "you're so silly"
const BadWinnerInputErrMsg = "invalid winner input, expect format of 'PlayerName wins'"

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	// 文字列としての標準入力から渡されるプレイヤー人数を数値に変換する
	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)

	// ２行目の標準入力を受け取る
	userInput := cli.readLine()
	winner, err := extractWinner(userInput)

	if err != nil {
		fmt.Fprint(cli.out, err.Error())
		return
	}

	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputErrMsg)
	}

	return strings.Replace(userInput, " wins", "", 1), nil
}

type Game interface {
	Start(numberOfPlayers int, alertDestination io.Writer)
	Finish(winner string)
}
