package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	// GetLeagueで最後まで読み込んでしまったので、カーソルの位置を先頭に戻した上でleagueを書き込む（なので、書き込み先の中身が全て書き換わる）
	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// 読み込み先の位置を先頭に戻しておく（VHSビデオのイメージ）
	f.database.Seek(0, io.SeekStart)
	// 読み込み
	league, _ := NewLeague(f.database)
	return league
}
