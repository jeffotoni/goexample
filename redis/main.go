// Back-End in Go server
// @jeffotoni
// 2019-04-08

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

var client *redis.Client

func NewClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Addr:     "redis:6379",
		//Addr:     "172.17.0.3:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return c
}

func SaveRedis(key_int int, value string) bool {
	key := strconv.Itoa(key_int)
	err := client.Set(key, value, 0).Err()
	if err != nil {
		log.Println("redis:: ", err)
		return false
	}

	return true
}

func main() {

	client = NewClient()

	pong, err := client.Ping().Result()
	if pong != "PONG" || err != nil {
		log.Println("Redis error: ", err)
		return
	}

	SaveRedis(1, "jefferson otoni")
	fmt.Println(pong)
}
