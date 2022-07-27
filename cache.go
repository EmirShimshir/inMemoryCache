package inMemoryCache

import (
	"fmt"
	"time"
)

type cacheMethods interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
	Delete(key string) error
}

type data struct {
	value      any
	timeDelete time.Time
}

type cache struct {
	storage map[string]data
}

func New() *cache {
	return &cache{
		storage: make(map[string]data),
	}
}

func (c *cache) Set(key string, value any, ttl time.Duration) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}

	if ttl == 0 {
		return fmt.Errorf("time-to-live equals zero")
	}

	c.storage[key] = data{value: value, timeDelete: time.Now().Add(ttl)}

	return nil
}

func (c *cache) Get(key string) (any, error) {
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

func (c *cache) Delete(key string) error {
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
