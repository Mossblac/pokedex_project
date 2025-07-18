package pokecache

import (
	"sync"
	"time"
)

/* Now, you’re adding a cache:
When you fetch a “page”
(i.e., make a request to a specific URL),
you store the response bytes using
the URL as the key.
If the user tries to go back (“mapb”),
you serve the data from the cache
if it’s there instead of making another
network call.
This cache also auto-deletes entries after
a certain time (the “reaper/cleanup”).*/

type Cache struct {
	Entry    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entry:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	valBytes, ok := c.Entry[key]
	if ok {
		return valBytes.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.Entry {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.Entry, k)

			}
		}
		c.mu.Unlock()
	}
}
