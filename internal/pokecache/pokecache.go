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

func NewCache(interval time.Duration) (c *Cache) {
	//create a cache struct, fill out Entry map
	//set the interval, start background tasks:
	//reaploop
}

func (*Cache) add(key string, val []byte) {
	//adds to cache
}
func (*Cache) get(key string) ([]byte, bool) {
	//obtains entry from cache
	//true if entry was found
	//fasle if not
}

func (*Cache) reapLoop() {
	//takes timer from NewCache as input
	//removes old data after a certain amount of time
}
