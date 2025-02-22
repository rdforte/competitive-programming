package lru

type LRU[K comparable, V any] struct {
	mp       map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
	capacity int
}

type Node[K comparable, V any] struct {
	key  K
	val  V
	prev *Node[K, V]
	next *Node[K, V]
}

func New[K comparable, V any](capacity int) *LRU[K, V] {
	head := &Node[K, V]{}
	tail := &Node[K, V]{}
	head.next, tail.prev = tail, head

	minCap := max(1, capacity)

	return &LRU[K, V]{
		mp:       make(map[K]*Node[K, V], minCap),
		head:     head,
		tail:     tail,
		capacity: minCap,
	}
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	n, ok := c.mp[key]
	if !ok {
		return *new(V), false
	}

	c.remove(n)
	c.insert(n)

	return n.val, true
}

func (c *LRU[K, V]) Set(key K, val V) {
	n, ok := c.mp[key]
	if ok {
		n.val = val
		c.remove(n)
		c.insert(n)
		return
	}

	if len(c.mp) == c.capacity {
		lastNode := c.tail.prev
		c.remove(lastNode)
		delete(c.mp, lastNode.key)
	}

	n = &Node[K, V]{
		key: key,
		val: val,
	}

	c.mp[key] = n
	c.insert(n)
}

func (c *LRU[K, V]) remove(n *Node[K, V]) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (c *LRU[K, V]) insert(n *Node[K, V]) {
	n.next = c.head.next
	n.next.prev = n
	c.head.next = n
	n.prev = c.head
}
