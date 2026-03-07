package store

import (
	"sync"
	"time"

	"app/internal/model"
)

type RenderConfigStore struct {
	mu     sync.RWMutex
	config map[string]model.RenderConfig
}

func NewRenderConfigStore() *RenderConfigStore {
	return &RenderConfigStore{
		config: make(map[string]model.RenderConfig),
	}
}

func (r *RenderConfigStore) Get(id string) (model.RenderConfig, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	config, exists := r.config[id]
	return config, exists
}

func (r *RenderConfigStore) Set(id string, config model.RenderConfig) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.config[id] = config
}

func (r *RenderConfigStore) Cleanup() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, config := range r.config {
		if config.HasExpired() {
			delete(r.config, id)
		}
	}
}

func (r *RenderConfigStore) StartCleanupThread() {
	go func() {
		for {
			r.Cleanup()
			// sleep for 1 hour
			time.Sleep(1 * time.Hour)
		}
	}()
}
