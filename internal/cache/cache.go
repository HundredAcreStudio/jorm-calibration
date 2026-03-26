package cache

import (
	"time"
)

// Entry is a cached value with an expiration time.
type Entry struct {
	Value     any
	ExpiresAt time.Time
}

// Cache is a simple in-memory cache. It is NOT safe for concurrent use.
// BUG: concurrent reads and writes will cause a data race.
type Cache struct {
	items map[string]Entry
	ttl   time.Duration
}

// New creates a cache with the given default TTL.
func New(ttl time.Duration) *Cache {
	return &Cache{
		items: make(map[string]Entry),
		ttl:   ttl,
	}
}

// Get returns the cached value for key, or nil if not found or expired.
func (c *Cache) Get(key string) any {
	entry, ok := c.items[key]
	if !ok {
		return nil
	}
	if time.Now().After(entry.ExpiresAt) {
		delete(c.items, key)
		return nil
	}
	return entry.Value
}

// Set stores a value in the cache with the default TTL.
func (c *Cache) Set(key string, value any) {
	c.items[key] = Entry{
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
	}
}

// Delete removes a key from the cache.
func (c *Cache) Delete(key string) {
	delete(c.items, key)
}

// Invalidate removes all expired entries from the cache.
func (c *Cache) Invalidate() {
	now := time.Now()
	for k, entry := range c.items {
		if now.After(entry.ExpiresAt) {
			delete(c.items, k)
		}
	}
}
