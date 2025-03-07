package main

import (
	"fmt"
	"math"
)

func main() {
	a := Constructor()

	a.Inc("hello")
	a.Inc("hello")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())
	a.Inc("leet")
	fmt.Println(a.GetMaxKey())
	fmt.Println(a.GetMinKey())

	fmt.Println("-------------------------")
	fmt.Println("-------------------------")

	b := Constructor()
	b.Inc("hello")   // 1
	b.Inc("goodbye") // 1
	b.Inc("hello")   // 2
	b.Inc("hello")   // 3
	fmt.Println(b.GetMaxKey())
	b.Inc("leet")  // 1
	b.Inc("code")  // 1
	b.Inc("leet")  // 2
	b.Dec("hello") // 2
	b.Inc("leet")  // 3
	b.Inc("code")  // 2
	b.Inc("code")  // 3
	fmt.Println(b.GetMaxKey())

	fmt.Println("-------------------------")
	fmt.Println("-------------------------")

	c := Constructor()
	c.Inc("hello") // 1
	c.Inc("hello") // 2
	fmt.Println(c.GetMaxKey())
	fmt.Println(c.GetMinKey())
	c.Dec("hello") // 1
	c.Dec("hello") // 0
	fmt.Println(c.GetMaxKey())
	c.Inc("hello") // 1
	fmt.Println(c.GetMinKey())

	fmt.Println("-------------------------")
	fmt.Println("-------------------------")

	d := Constructor()
	d.Inc("a") // 1
	d.Inc("b") // 1
	d.Inc("b") // 2
	d.Inc("c") // 1
	d.Inc("c") // 2
	d.Inc("c") // 3
	d.Dec("b") // 1
	d.Dec("b") // 0
	fmt.Println(d.GetMinKey())
	d.Dec("a") // 0
	fmt.Println(d.GetMaxKey())
	fmt.Println(d.GetMinKey())
}

type AllOne struct {
	mp      map[string]*Node
	minHead *Node
	maxTail *Node
}

func Constructor() AllOne {
	minHead := &Node{val: 0, mp: make(map[string]struct{})}
	maxTail := &Node{val: math.MaxInt, mp: make(map[string]struct{})}

	minHead.next, maxTail.prev = maxTail, minHead

	return AllOne{
		mp:      make(map[string]*Node),
		minHead: minHead,
		maxTail: maxTail,
	}
}

func (a *AllOne) Inc(key string) {
	n, ok := a.mp[key]

	if !ok {
		if a.minHead.next.val != a.minHead.val+1 {
			newNode := &Node{val: 1, mp: make(map[string]struct{})}
			a.insert(a.minHead, newNode)
		}
		a.minHead.next.mp[key] = struct{}{}
		a.mp[key] = a.minHead.next
		return
	}

	if n.next.val != n.val+1 {
		newNode := &Node{val: n.val + 1, mp: make(map[string]struct{})}
		a.insert(n, newNode)
	}

	delete(n.mp, key)
	n.next.mp[key] = struct{}{}
	a.mp[key] = n.next

	if len(n.mp) == 0 {
		a.remove(n)
	}
}

func (a *AllOne) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (a *AllOne) insert(prevNode, newNode *Node) {
	prevNode.next.prev, newNode.next = newNode, prevNode.next
	prevNode.next, newNode.prev = newNode, prevNode
}

func (a *AllOne) Dec(key string) {
	n, ok := a.mp[key]
	if !ok {
		return
	}

	if n.prev.val != n.val-1 {
		newNode := &Node{val: n.val - 1, mp: make(map[string]struct{})}
		a.insert(n.prev, newNode)
	}

	delete(n.mp, key)

	// if start node
	if n.prev.val == 0 {
		delete(a.mp, key)
		if len(n.mp) == 0 {
			a.remove(n)
		}
		return
	}

	a.mp[key] = n.prev
	n.prev.mp[key] = struct{}{}

	if len(n.mp) == 0 {
		a.remove(n)
	}
}

func (a *AllOne) GetMaxKey() string {
	for key := range a.maxTail.prev.mp {
		return key
	}

	return ""
}

func (a *AllOne) GetMinKey() string {
	for key := range a.minHead.next.mp {
		return key
	}

	return ""
}

type Node struct {
	val  int
	mp   map[string]struct{}
	next *Node
	prev *Node
}
