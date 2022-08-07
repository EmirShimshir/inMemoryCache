package inMemoryCache

import (
	"fmt"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
}

type cacheData struct {
	value      any
	timeDelete time.Time
}

type cacheMem struct {
	storage map[string]cacheData
	mu      sync.RWMutex
}

func New() Cache {
	return &cacheMem{
		storage: make(map[string]cacheData),
	}
}

func (c *cacheMem) Set(key string, value any, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		return fmt.Errorf("empty key")
	}

	if ttl == 0 {
		return fmt.Errorf("time-to-live equals zero")
	}

	c.storage[key] = cacheData{value: value, timeDelete: time.Now().Add(ttl)}

	return nil
}

func (c *cacheMem) Get(key string) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		return nil, fmt.Errorf("empty key")
	}

	data, exists := c.storage[key]

	if !exists {
		return nil, fmt.Errorf("no value for the key %s", key)
	}

	if time.Now().After(data.timeDelete) {
		delete(c.storage, key)

		return nil, fmt.Errorf("no value for the key %s", key)
	}

	return data.value, nil
}

func (c *cacheMem) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if key == "" {
		return fmt.Errorf("empty key")
	}

	_, exists := c.storage[key]

	if !exists {
		return fmt.Errorf("no value for the key %s", key)
	}

	delete(c.storage, key)

	return nil
}
