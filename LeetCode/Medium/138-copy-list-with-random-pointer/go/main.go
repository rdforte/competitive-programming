package main

import "fmt"

func main() {

	three := &Node{3, nil, nil}
	two := &Node{2, three, nil}
	head := &Node{1, two, nil}
	two.Random = head

	h := copyRandomList(head)

	fmt.Println(h)
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)

	og := head
	c := &Node{-1, nil, nil}
	newHead := c
	for og != nil {
		c.Next = &Node{og.Val, nil, nil}
		c = c.Next
		mp[og] = c
		og = og.Next
	}

	c = newHead.Next
	for head != nil {
		if val, ok := mp[head.Random]; ok {
			c.Random = val
		}
		c = c.Next
		head = head.Next
	}

	return newHead.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
