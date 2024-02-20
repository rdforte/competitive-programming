package main

import "fmt"

func main() {
	q := createQueue[int]()
	q.Enqueue(1)
	fmt.Println(q.Peek())
	q.Enqueue(2)
	fmt.Println(q.Peek())
	q.Enqueue(3)
	fmt.Println(q.Peek())
	q.Deque()
	fmt.Println(q.Len())
	q.Deque()
	fmt.Println(q.Len())
}

func createQueue[T any]() queue[T] {
	return queue[T]{}
}

type queue[T any] struct {
	len  int
	head *ListNode[T]
	tail *ListNode[T]
}

func (q *queue[T]) Len() int {
	return q.len
}

func (q *queue[T]) Enqueue(item T) {
	node := &ListNode[T]{val: item}
	q.len++
	if q.tail == nil {
		q.tail = node
		q.head = q.tail
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *queue[T]) Deque() (T, error) {
	if q.head == nil {
		return *new(T), &QueueError{"can not deque empty queue"}
	}
	q.len--
	head := q.head
	q.head = head.next

	if q.len == 0 {
		q.tail = nil
	}

	return head.val, nil
}

func (q *queue[T]) Peek() (T, error) {
	if q.head == nil {
		return *new(T), &QueueError{"can not peek empty queue"}
	}
	return q.head.val, nil
}

type ListNode[T any] struct {
	val  T
	next *ListNode[T]
}

type QueueError struct {
	message string
}

func (s *QueueError) Error() string {
	return s.message
}
