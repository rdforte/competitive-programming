package main

import "fmt"

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
	mp       map[int]*node
	head     *node
	tail     *node
}

func Constructor(capacity int) *LRUCache {
	head := &node{}
	tail := &node{}
	// All the nodes that we insert should be between the head and tail nodes.
	// This will resolve the issue of calling next/prev on a nil node and causing a panic.
	head.next = tail
	tail.prev = head
	return &LRUCache{capacity, make(map[int]*node, capacity), head, tail}
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.mp[key]
	if !ok {
		return -1
	}
	c.remove(n)
	c.add(n)
	return n.val
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.mp[key]; ok {
		c.remove(n)
		n.val = value
		c.add(n)
		return
	}

	if len(c.mp) == c.capacity {
		// evict tail to make room for new key
		delete(c.mp, c.tail.prev.key)
		c.remove(c.tail.prev)
	}

	n := &node{val: value, key: key}
	c.add(n)
	c.mp[key] = n
}

func (c *LRUCache) remove(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) add(node *node) {
	c.head.next.prev = node
	node.next = c.head.next
	c.head.next = node
	node.prev = c.head
}

type node struct {
	val  int
	key  int
	next *node
	prev *node
}
