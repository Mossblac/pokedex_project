package pokecache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// t.Errorf("describes error and fails test")
func TestReapingLoop(t *testing.T) {
	cache := NewCache(5 * time.Millisecond)
	key := "https://testURL.com"
	val := []byte("testbytes")
	cache.Add(key, val)

	output, ok := cache.Get(key)
	assert.True(t, ok, "entry should exist right after insertion")
	assert.Equal(t, val, output, "the output should be what was added")

	time.Sleep(10 * time.Millisecond)
	_, ok = cache.Get(key)
	assert.False(t, ok, "entry should have been reaped after interval")

}
