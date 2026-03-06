package store

import (
	"fmt"
	"log"
	"sync"
	"time"

	"app/internal/model"
)

type NamesStore struct {
	mu    sync.RWMutex
	names map[string][]model.Name
}

func (n *NamesStore) Cleanup() {
	log.Println("Cleaning up expired names...")
	cleanedCount := 0
	n.mu.Lock()
	defer n.mu.Unlock()
	for id, names := range n.names {
		var validNames []model.Name
		for _, name := range names {
			if !name.HasExpired() {
				validNames = append(validNames, name)
			} else {
				cleanedCount++
			}
		}
		n.names[id] = validNames
	}
	log.Printf("Cleanup complete. Removed %d expired names.\n", cleanedCount)
}

func (n *NamesStore) StartCleanupThread() {
	go func() {
		log.Println("Starting cleanup thread for expired names...")
		for {
			n.Cleanup()
			// sleep for 1 hour
			time.Sleep(1 * time.Hour)
			log.Printf("Next cleanup will run at: %s\n", time.Now().Add(1*time.Hour).Format(time.RFC1123))
		}
	}()
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
	// prepend name to the list of names for the id
	n.names[id] = append([]model.Name{{Name: name}}, n.names[id]...)
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
