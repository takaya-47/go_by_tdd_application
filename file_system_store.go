package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	}

	// GetLeagueで最後まで読み込んでしまったので、カーソルの位置を先頭に戻した上でleagueを書き込む（なので、書き込み先の中身が全て書き換わる）
	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) GetLeague() League {
	// 読み込み先の位置を先頭に戻しておく（VHSビデオのイメージ）
	f.database.Seek(0, io.SeekStart)
	// 読み込み
	league, _ := NewLeague(f.database)
	return league
}