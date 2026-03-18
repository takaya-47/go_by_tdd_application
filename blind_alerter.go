package poker

import (
	"fmt"
	"os"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

type BlindAlerterFunc func(duration time.Duration, amount int)

func (f BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	f(duration, amount)
}

func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}

// TODO: https://andmorefine.gitbook.io/learn-go-with-tests/build-an-application/time#nitesutowoku-2

// ブランチ：feature/time/test_2

// 次の「最初にテストを書く」まで進んだらPRを作成し、CI実行を確認すること