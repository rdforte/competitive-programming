package lru_test

import (
	"fmt"
	"testing"

	"github.com/rdforte/competitive-programming/random/lru"
)

func TestLRU(t *testing.T) {
	cache := lru.New[string, string](2)

	cache.Set("ryan", "ryan")
	fmt.Println(cache.Get("ryan"))
	fmt.Println("-----------------")
	cache.Set("bianca", "bianca")
	fmt.Println(cache.Get("bianca"))
	fmt.Println(cache.Get("ryan"))
	fmt.Println("-----------------")
	cache.Set("shernaz", "shernaz")
	fmt.Println(cache.Get("shernaz"))
	fmt.Println(cache.Get("bianca"))
	fmt.Println(cache.Get("ryan"))
}
