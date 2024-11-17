package main

import "fmt"

func main() {
	test1()
	fmt.Println("---------------------")
	test2()
	fmt.Println("---------------------")
	test3()
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	next := dummy

	for l1 != nil && l2 != nil {
		next.Next = &ListNode{}
		next = next.Next
		sum := l1.Val + l2.Val + carry
		if sum/10 >= 1 {
			carry = 1
			sum %= 10
		} else {
			carry = 0
		}
		next.Val = sum
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		next.Next = &ListNode{}
		next = next.Next
		sum := l1.Val + carry
		if sum/10 >= 1 {
			carry = 1
			sum %= 10
		} else {
			carry = 0
		}
		next.Val = sum
		l1 = l1.Next
	}

	for l2 != nil {
		next.Next = &ListNode{}
		next = next.Next
		sum := l2.Val + carry
		if sum/10 >= 1 {
			carry = 1
			sum %= 10
		} else {
			carry = 0
		}
		next.Val = sum
		l2 = l2.Next
	}

	if carry == 1 {
		next.Next = &ListNode{}
		next = next.Next
		next.Val = 1
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func test1() {
	l1Node3 := &ListNode{3, nil}
	l1Node2 := &ListNode{4, l1Node3}
	l1Node1 := &ListNode{2, l1Node2}

	l2Node3 := &ListNode{4, nil}
	l2Node2 := &ListNode{6, l2Node3}
	l2Node1 := &ListNode{5, l2Node2}

	res := addTwoNumbers(l1Node1, l2Node1)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

func test2() {
	l1Node7 := &ListNode{9, nil}
	l1Node6 := &ListNode{9, l1Node7}
	l1Node5 := &ListNode{9, l1Node6}
	l1Node4 := &ListNode{9, l1Node5}
	l1Node3 := &ListNode{9, l1Node4}
	l1Node2 := &ListNode{9, l1Node3}
	l1Node1 := &ListNode{9, l1Node2}

	l2Node4 := &ListNode{9, nil}
	l2Node3 := &ListNode{9, l2Node4}
	l2Node2 := &ListNode{9, l2Node3}
	l2Node1 := &ListNode{9, l2Node2}

	res := addTwoNumbers(l1Node1, l2Node1)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}

func test3() {
	l1Node1 := &ListNode{0, nil}
	l2Node1 := &ListNode{0, nil}

	res := addTwoNumbers(l1Node1, l2Node1)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
