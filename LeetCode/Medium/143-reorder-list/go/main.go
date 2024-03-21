package main

import "fmt"

func main() {

	five := &ListNode{5, nil}
	four := &ListNode{4, five}
	three := &ListNode{3, four}
	two := &ListNode{2, three}
	head := &ListNode{1, two}

	reorderList(head)
	fmt.Println(head)
}

func reorderList(head *ListNode) {
	mid := head
	fast := head

	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}

	// reverse mid
	var prev *ListNode
	next := mid
	for next != nil {
		cur := next
		next = cur.Next
		cur.Next = prev
		prev = cur
	}

	primary := head
	secondary := prev

	// The secondary chains length is less than the primary chains length.
	// Because the last item of secondary is the same as the last item of primary we disregard secondarys last node.
	for secondary.Next != nil {
		tmp := primary.Next
		primary.Next = secondary
		secondary = secondary.Next
		primary.Next.Next = tmp
		primary = tmp
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
