package main

import (
	"container/list"
	"fmt"
)

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

// Building a LRU CACHE.
// Cache should be fast so we want O(1) to store and get a value from the cache.
// We also want the ability to remove the last rescently used item from the cache.
// maps are O(1) lookup time and we can easily evict something from the cache by using
// a linked list where by we can remove the tail node. If we used an array instead of a linked
// list this this would be consecutive memory which means we would have to shuffle all the elements
// in the array which would be O(n) time.

type LRUCache struct {
	capacity int
	mp       map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		mp:       make(map[int]*list.Element, capacity),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.mp[key]; ok {
		c.list.MoveToFront(n)
		return n.Value.(node).value
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.mp[key]; ok {
		c.list.MoveToFront(n)
		n.Value = node{key, value}
		return
	}

	if len(c.mp) == c.capacity {
		last := c.list.Back().Value
		delete(c.mp, last.(node).key)
		c.list.Remove(c.list.Back())
	}

	n := c.list.PushFront(node{key, value})
	c.mp[key] = n
}

type node struct {
	key   int
	value int
}
