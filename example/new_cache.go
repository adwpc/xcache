package main

import (
	"fmt"
	"time"

	"github.com/adwpc/xcache"
	"github.com/allegro/bigcache/v3"
)

type MyCache struct {
	xc *xcache.XCache
}

func NewMyCache() *MyCache {
	config := bigcache.Config{
		Shards:           1024,             // shards number
		LifeWindow:       time.Hour,        // item life time is 1 Hour
		CleanWindow:      time.Second * 10, // clean cycle is 10s
		HardMaxCacheSize: 1000,             // max cache size is 1000MB
		OnRemove: func(key string, entry []byte) {
			fmt.Printf("delete key:%s %v\n", key, entry)
		},
	}
	return &MyCache{xc: xcache.New(config)}
}

func (m *MyCache) Set(key, value interface{}) {
	xcache.XSet(m.xc, key, value)
}

// Get will get a interface, and you need to convert it by ".(type)"
func (m *MyCache) Get(key interface{}) interface{} {
	return xcache.XGet[interface{}, interface{}](m.xc, key)
}

func (m *MyCache) SetStr(key, value string) {
	xcache.XSet(m.xc, key, value)
}

func (m *MyCache) GetStr(key string) string {
	return xcache.XGet[string, string](m.xc, key)
}

func (m *MyCache) SetStrFloat(key string, value float64) {
	xcache.XSet(m.xc, key, value)
}

func (m *MyCache) GetFloatByStr(key string) float64 {
	return xcache.XGet[string, float64](m.xc, key)
}

// this is a custom get method, otherwise Get will get a map[Str:hello] val
func (m *MyCache) GetJson(key interface{}) MyJson {
	return xcache.XGet[interface{}, MyJson](m.xc, key)
}

type MyJson struct {
	Str string `json: str` //must be UpperCase
}

// this is a new cache, custom usage example
func main() {
	mycache := NewMyCache()
	mycache.Set("a", "b")
	mycache.Set(1.0, 2.0)

	fmt.Printf("%+v\n", mycache.Get("a").(string))
	fmt.Printf("%+v\n", mycache.Get(1.0).(float64))

	mycache.SetStr("c", "d")
	fmt.Printf("%+v\n", mycache.GetStr("c"))

	mycache.SetStrFloat("e", 0.01)
	fmt.Printf("%+v\n", mycache.GetFloatByStr("e"))

	mycache.Set("json", MyJson{"hello"})
	val := mycache.GetJson("json")
	fmt.Printf("val=%+v\n", val)

}
