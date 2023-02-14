package main

import (
	"github.com/patrickmn/go-cache"

	"time"
)

var c = cache.New(10*time.Minute, 10*time.Minute)

func main() {

}

func Set(key string, val int, ttl time.Duration) {
	c.Set(key, val, ttl)
}

func Get(key string) interface{} {
	accessTokenC, _ := c.Get(key)
	return accessTokenC
}
