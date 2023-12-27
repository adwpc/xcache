package xcache

import (
	"context"
	"fmt"

	"github.com/allegro/bigcache/v3"
)

// XCache is a wrapper of bigcache.BigCache
type XCache struct {
	cache *bigcache.BigCache
}

// New returns a new XCache
func New(config bigcache.Config) *XCache {
	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &XCache{
		cache: cache,
	}
}

func XGet[X, Y any](xc *XCache, key X) Y {
	return get[X, Y](xc.cache, key) //cannot infer Y
}

func XSet[X, Y any](xc *XCache, key X, value Y) {
	set(xc.cache, key, value)
}
