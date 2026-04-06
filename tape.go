package poker

import (
	"io"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	// ファイルサイズを0（中身の全削除）にする
	t.file.Truncate(0)
	// 書き込み位置を先頭に戻す
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
