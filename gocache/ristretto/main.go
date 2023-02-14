package main

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

var c, err = ristretto.NewCache(&ristretto.Config{
	NumCounters: 1e7,     // number of keys to track frequency of (10M).
	MaxCost:     1 << 30, // maximum cost of cache (1GB).
	BufferItems: 64,      // number of keys per Get buffer.
})

func main() {

}

func Set(key string, val int, ttl time.Duration) {
	c.Set(key, val, int64(ttl))
}

func Get(key string) interface{} {
	accessTokenC, _ := c.Get(key)
	return accessTokenC
}
