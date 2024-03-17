package main

import "fmt"

func main() {
	list1Node3 := ListNode{4, nil}
	list1Node2 := ListNode{2, &list1Node3}
	list1Root := ListNode{1, &list1Node2}
	list2Node3 := ListNode{4, nil}
	list2Node2 := ListNode{3, &list2Node3}
	list2Root := ListNode{1, &list2Node2}

	head := mergeTwoLists(&list1Root, &list2Root)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	next := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next.Next = list1
			list1 = list1.Next
		} else {
			next.Next = list2
			list2 = list2.Next
		}
		next = next.Next
	}

	if list1 != nil {
		next.Next = list1
	}

	if list2 != nil {
		next.Next = list2
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
