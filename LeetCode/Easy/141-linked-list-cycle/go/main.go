package main

import "fmt"

func main() {
	test1()
	fmt.Println("--------")
	test2()
	fmt.Println("--------")
	test3()
	fmt.Println("--------")
	test4()
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func test1() {
	n4 := &ListNode{-4, nil}
	n3 := &ListNode{0, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{3, n2}
	n4.Next = n2

	fmt.Println(hasCycle(n1))
}

func test2() {
	n2 := &ListNode{2, nil}
	n1 := &ListNode{1, n2}
	n2.Next = n1

	fmt.Println(hasCycle(n1))
}

func test3() {
	n1 := &ListNode{1, nil}
	fmt.Println(hasCycle(n1))
}

func test4() {
	fmt.Println(hasCycle(nil))
}
