package database

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})
}

// jika arraw ada di sebelah kiri maka channel tersebut akan
func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {

		for {
			time.Sleep(3 * time.Second)

			key := <-ch

			Cache.Del(context.Background(), key)

			fmt.Println("cache clear..." + key)

		}

	}(CacheChannel)
}

// jika arrow ada di sebelah kanan maka channel tersebut akan menerima value key
func ClearCache(keys ...string) {
	for _, key := range keys {

		CacheChannel <- key
	}
}
