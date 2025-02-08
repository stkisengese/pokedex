package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cachedData   map[string]cacheEntry // store cached data
	mutex        sync.RWMutex          // To ensure thread-safety for concurrent access to the map
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time // timestamp of entry creation
	val       []byte    // cached data
}

// NewCache creates a new cache with a configurable reaping interval.
func NewCache(reapInterval time.Duration) *Cache {
	c := &Cache{
		cachedData:   make(map[string]cacheEntry),
		reapInterval: reapInterval,
	}
	go c.reapLoop() // start reaping loop in a goroutine
	return c
}

// Add adds a new entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cachedData[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves an entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, exists := c.cachedData[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

// reapLoop runs in a gorutine to periodically remove the expired entries.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.cachedData {
			if time.Since(entry.createdAt) > c.reapInterval {
				delete(c.cachedData, key)
			}
		}
		c.mutex.Unlock()
	}
}
