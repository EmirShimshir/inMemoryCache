package inMemoryCache

import (
	"fmt"
	"time"
)

type CacheMethods interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
}

type data struct {
	value      any
	timeDelete time.Time
}

type Cache struct {
	storage map[string]data
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]data),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}

	if ttl == 0 {
		return fmt.Errorf("time-to-live equals zero")
	}

	c.storage[key] = data{value: value, timeDelete: time.Now().Add(ttl)}

	return nil
}

func (c *Cache) Get(key string) (any, error) {
	if key == "" {
		return nil, fmt.Errorf("empty key")
	}

	data, exists := c.storage[key]

	if !exists {
		return nil, fmt.Errorf("no value for the key %s", key)
	}

	if time.Now().After(data.timeDelete) {
		_ = c.Delete(key)

		return nil, fmt.Errorf("no value for the key %s", key)
	}

	return data.value, nil
}

func (c *Cache) Delete(key string) error {
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
