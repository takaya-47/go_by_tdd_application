package poker

import "io"

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin("Chris") // 固定のプレイヤー名。あとで標準入力から受け取った名前にする。
}