package main

import "fmt"

func main() {
	five := &ListNode{5, nil}
	four := &ListNode{4, five}
	three := &ListNode{3, four}
	two := &ListNode{2, three}
	head := &ListNode{1, two}

	h := removeNthNodeFromEnd(head, 5)

	fmt.Println(h)
}

func removeNthNodeFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	slow := dummy
	fast := dummy

	// iterate n + 1 hence why use <=n and not just < n. This is because we want the node before the one we are removing.
	// we need the dummy node just in case we decide to remove the very first node.
	// if you think about it if we seperate the pointers by n ie n = 2 then by the time the fast pointer hits the end our slow pointer will be 2 behind.
	for i := 0; i <= n && fast != nil; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	cur := slow
	for i := 0; i < 2 && slow != nil; i++ {
		slow = slow.Next
	}

	cur.Next = slow

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
