package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items map[string]Item
	sync.RWMutex
}

type Item struct {
	Value      interface{}
	CreateTime time.Time
	Expiration int64
}

func New() *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	expiration := time.Now().Add(ttl).UnixNano()

	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		CreateTime: time.Now(),
		Expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RLock()

	item, exists := c.items[key]
	if !exists {
		return nil, errors.New("cache element not found")
	}
	return item.Value, nil
}

func (c *Cache) Delete(key string) error {
	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return nil
}
