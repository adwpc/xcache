package xcache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/allegro/bigcache/v3"
)

var (
	cache *bigcache.BigCache
)

func init() {
	Init(time.Hour)
}

// Init initializes the cache
func Init(expire time.Duration) {
	fmt.Println("init cache")
	// config: https://pkg.go.dev/github.com/allegro/bigcache/v3#Config
	cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(expire))
}

func toBytes[T any](value T) []byte {
	bytes, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("error converting value to bytes: %v", err)
	}
	return bytes
}

func set[X, Y any](c *bigcache.BigCache, key X, value Y) {
	if c == nil {
		return
	}

	c.Set(fmt.Sprintf("%v", key), toBytes(value))
}

func get[X, Y any](c *bigcache.BigCache, key X) Y {
	var value Y
	if c == nil {
		fmt.Println("cache is nil")
		return value
	}

	bytes, err := c.Get(fmt.Sprintf("%v", key))
	if err != nil {
		fmt.Println(err)
		return value
	}

	err = json.Unmarshal(bytes, &value)
	if err != nil {
		log.Fatalf("error converting bytes to value: %v", err)
	}
	return value
}

func Get[X, Y any](key X) Y {
	return get[X, Y](cache, key) //cannot infer Y
}

func Set[X, Y any](key X, value Y) {
	set(cache, key, value)
}
