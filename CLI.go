package poker

type CLI struct {
	store PlayerStore
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin("Cleo") // 固定のプレイヤー名。あとで標準入力から受け取った名前にする。
}