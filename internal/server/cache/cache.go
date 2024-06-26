package cache

import (
	"L0/internal/bd"
	"L0/internal/structs"
	"errors"
	"log"
	"sync"
)

var Cache *RwCache

type RwCache struct {
	mu    sync.RWMutex
	cache map[string]*structs.Order
}

func NewRwCache() *RwCache {
	return &RwCache{cache: make(map[string]*structs.Order)}
}

func (c *RwCache) Add(key string, value *structs.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *RwCache) Get(key string) (*structs.Order, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.cache[key]
	if !exists {
		return nil, errors.New("not found")
	}
	return value, nil
}

func (c *RwCache) IsExist(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, exists := c.cache[key]
	return exists
}

func (c *RwCache) Restore() error {
	arr, err := bd.GetUIDs()
	if err != nil {
		log.Printf("Faild to restore data\n")
		return err
	}
	for _, data := range arr {
		c.cache[data.OrderUID] = data
	}
	return nil
}
