package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// 読み込み先の位置を先頭に戻しておく（VHSビデオのイメージ）
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}