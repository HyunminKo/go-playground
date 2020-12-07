package main

import (
	"fmt"

	"github.com/muesli/cache2go"
)

func main() {
	key := "SameKey"
	tct := cache2go.Cache("TestCacheTable")
	tct.Add(key, 0, "raw data")
	item, err := tct.Value(key)
	if err != nil {
		fmt.Println("The key does not exist", key)
		return
	}
	fmt.Println("The data is", item.Data())

	tct.Add(key, 0, "raw data2")
	item, err = tct.Value(key)
	if err != nil {
		fmt.Println("The key does not exist", key)
		return
	}
	fmt.Println("The data is", item.Data())
}
