package store

import (
	"fmt"
	"sync"

	"app/internal/model"
)

type NamesStore struct {
	mu    sync.RWMutex
	names map[string][]model.Name
}

func NewNamesStore() *NamesStore {
	return &NamesStore{
		names: make(map[string][]model.Name),
	}
}

func (n *NamesStore) Create(id string, name string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	fmt.Printf("Creating name for id: %s\n", id)
	fmt.Println(n.names[id])
	n.names[id] = append(n.names[id], model.Name{Name: name})
	fmt.Println(n.names[id])
}

func (n *NamesStore) List(id string) []model.Name {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if name, ok := n.names[id]; ok {
		return name
	}
	return []model.Name{}
}
