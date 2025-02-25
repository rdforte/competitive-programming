package main

import "math"

func main() {
}

type AllOne struct {
	mp      map[string]*Node
	minHead *Node
	maxTail *Node
}

func Constructor() AllOne {
	minHead := &Node{val: 0}
	maxTail := &Node{val: math.MaxInt}

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
	}

	if n.next.val == math.MaxInt {
		// insert new node
		newNode := &Node{val: n.val + 1, mp: map[string]struct{}{key: {}}}
		a.insert(n, newNode)
		a.mp[key] = newNode
		return
	}

	n.next.mp[key] = struct{}{}
	a.mp[key] = n.next
}

func (a *AllOne) insert(prevNode, newNode *Node) {
	prevNode.next.prev, newNode.next = newNode, prevNode.next
	prevNode.next, newNode.prev = newNode, prevNode
}

func (a *AllOne) Dec(key string) {
}

func (a *AllOne) GetMaxKey() string {
}

func (a *AllOne) GetMinKey() string {
}

type Node struct {
	val  int
	mp   map[string]struct{}
	next *Node
	prev *Node
}
