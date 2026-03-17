package poker

import (
	"io"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()

	tape := &tape{file}
	tape.Write([]byte("abc")) // "abc"を[]byteにキャスト

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)
	got := string(newFileContents) // []byteをstringにキャスト
	want := "abc"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}