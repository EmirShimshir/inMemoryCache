package inMemoryCache

import "fmt"

type CacheMethods interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type Cache struct {
	storage map[string]interface{}
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) error {
	if key == "" {
		return fmt.Errorf("empty key")
	}

	c.storage[key] = value

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	if key == "" {
		return nil, fmt.Errorf("empty key")
	}

	value, exists := c.storage[key]

	if !exists {
		return nil, fmt.Errorf("no value for the key %s", key)
	}

	return value, nil
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
