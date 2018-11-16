package cache

import (
	c "github.com/patrickmn/go-cache"
)

// Cache interface
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Keys() []string
}

type cache struct {
	cache *c.Cache
}

// New returns a cache implementation
func New() Cache {
	return &cache{
		cache: c.New(-1, -1),
	}
}

// Get returns item if exists
func (c *cache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

// Set add/update item on cache
func (c *cache) Set(key string, value interface{}) {
	c.cache.Set(key, value, -1)
}

// Keys return a list of keys from cache
func (c *cache) Keys() []string {
	items := c.cache.Items()
	keys := []string{}

	for key := range items {
		keys = append(keys, key)
	}

	return keys
}
