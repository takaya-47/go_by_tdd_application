package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(r io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

type League []Player

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
