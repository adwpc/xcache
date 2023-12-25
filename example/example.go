package main

import (
	"fmt"

	"github.com/adwpc/xcache"
)

func main() {
	// xcache.Init(time.Hour*12) // not nessary, will be called automatically(expire is 1 Hour)
	xcache.Set("a", "b")
	xcache.Set(1.0, 2)
	fmt.Println(xcache.Get[string, string]("a"))
	fmt.Println(xcache.Get[float64, int](1.0))

	type json struct {
		Str string //must be UpperCase
	}
	xcache.Set("json", json{"hello"})
	fmt.Println(xcache.Get[string, json]("json"))

}
