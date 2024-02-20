package main

import "fmt"

func main() {
	s := createStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Printf("stack len: %d \n", s.Len())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Printf("stack len: %d \n", s.Len())
	fmt.Println(s.Peek())
}

func createStack[T any]() stack[T] {
	return stack[T]{}
}

type stack[T any] struct {
	len  int
	head *ListNode[T]
}

func (s *stack[T]) Peek() (T, error) {
	if s.len == 0 {
		return *new(T), &StackError{"cant peek empty stack"}
	}
	return s.head.val, nil
}

func (s *stack[T]) Push(item T) {
	s.len++
	newHead := ListNode[T]{item, s.head}
	s.head = &newHead
}

func (s *stack[T]) Pop() (T, error) {
	if s.len == 0 {
		return *new(T), &StackError{"cant pop off of empty stack"}
	}
	s.len--
	head := s.head
	s.head = head.prev
	return head.val, nil
}

func (s *stack[T]) Len() int {
	return s.len
}

type ListNode[T any] struct {
	val  T
	prev *ListNode[T]
}

type StackError struct {
	message string
}

func (s *StackError) Error() string {
	return s.message
}
