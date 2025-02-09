package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachedData   sync.Map // store cached data
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time // timestamp of entry creation
	val       []byte    // cached data
}

// NewCache creates a new cache with a configurable reaping interval.
func NewCache(reapInterval time.Duration) *Cache {
	c := &Cache{
		reapInterval: reapInterval,
	}
	go c.reapLoop() // start reaping loop in a goroutine
	return c
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.cachedData.Store(key, cacheEntry{
		createdAt: time.Now(),
		val:       val,
	})
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	if entryInterface, ok := c.cachedData.Load(key); ok {
		entry, ok := entryInterface.(cacheEntry) // Type assertion
		if !ok {
			return nil, false
		}
		return entry.val, true
	}
	return nil, false
}

// reapLoop runs in a gorutine to periodically remove the expired entries.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	defer ticker.Stop()

	for range ticker.C {
		keysToDelete := []string{}

		c.cachedData.Range(func(key, value interface{}) bool {
			entry, ok := value.(cacheEntry)
			if !ok {
				return true
			}
			if time.Since(entry.createdAt) > c.reapInterval {
				keysToDelete = append(keysToDelete, key.(string))
			}
			return true
		})
		for _, key := range keysToDelete {
			c.cachedData.Delete(key)
		}
	}
}
