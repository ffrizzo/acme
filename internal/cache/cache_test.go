package cache_test

import (
	"testing"

	"github.com/ffrizzo/acme/internal/cache"
)

func TestNewCache(t *testing.T) {
	var (
		key   = "test"
		value = "test@test"
	)
	cache := cache.New()

	cache.Set(key, value)
	v, ok := cache.Get(key)
	if !ok {
		t.Errorf("Key %s is not present on cache", key)
	}

	cacheValue, ok := v.(string)
	if !ok {
		t.Error("Value is not a string")
	}

	if cacheValue != value {
		t.Errorf("Value %s does not match with value sent to cache", cacheValue)
	}
}

func TestCacheUpdate(t *testing.T) {
	var (
		key   = "test"
		value = "test@test"
	)
	cache := cache.New()

	cache.Set(key, value)
	v, ok := cache.Get(key)
	if !ok {
		t.Errorf("Key %s is not present on cache", key)
	}

	cacheValue, ok := v.(string)
	if !ok {
		t.Error("Value is not a string")
	}

	if cacheValue != value {
		t.Errorf("Value %s does not match with value sent to cache", cacheValue)
	}

	value = "test@acme"
	cache.Set(key, value)
	v, ok = cache.Get(key)
	if !ok {
		t.Errorf("Key %s is not present on cache", key)
	}

	cacheValue, ok = v.(string)
	if !ok {
		t.Error("Value is not a string")
	}

	if cacheValue != value {
		t.Errorf("Value %s does not match with value sent to cache", cacheValue)
	}

}

func TestCacheKeys(t *testing.T) {
	cache := cache.New()

	cache.Set("1", "1")
	cache.Set("2", "2")

	keys := cache.Keys()
	if len(keys) != 2 {
		t.Errorf("Cache returns %d keys and %d is expected", len(keys), 2)
	}

	for _, k := range keys {
		_, ok := cache.Get(k)
		if !ok {
			t.Errorf("Value not found on cache for key %s", k)
		}
	}
}
