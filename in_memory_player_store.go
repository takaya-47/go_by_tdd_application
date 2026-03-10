package main

import "sync"

// PlayerStoreを実装するインメモリのストア
type InMemoryPlayerStore struct{
	store map[string]int
	lock  sync.Mutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()

	if score, ok := i.store[name]; ok {
		return score
	}
	return 0
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return nil
}