package main

type InMemoryPlayerStore struct{
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	if score, ok := i.store[name]; ok {
		return score
	}
	return 0
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}