package main

import "fmt"

func main() {
	node3 := ListNode[int]{3, nil}
	node2 := ListNode[int]{2, &node3}
	node1 := ListNode[int]{1, &node2}

	for i := &node1; i != nil; i = i.next {
		fmt.Printf("%v, ", i.val)
	}
}

type ListNode[T any] struct {
	val  T
	next *ListNode[T]
}
