package main

import "container/heap"

func main() {
}

type StockPrice struct {
	minHeap    *MinHeap
	maxHeap    *MaxHeap
	prices     map[int]int
	latestTime int
}

func Constructor() StockPrice {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	return StockPrice{
		minHeap: minHeap,
		maxHeap: maxHeap,
		prices:  make(map[int]int),
	}
}

// Pushing to the heap requires log.N time so for N updates this requires N.log.N time.
func (s *StockPrice) Update(timestamp int, price int) {
	s.prices[timestamp] = price

	p := Price{time: timestamp, value: price}
	heap.Push(s.minHeap, p)
	heap.Push(s.maxHeap, p)

	s.latestTime = max(s.latestTime, timestamp)
}

func (s *StockPrice) Current() int {
	return s.prices[s.latestTime]
}

// NOTE - with max and min heaps the top element is always index 0

func (s *StockPrice) Maximum() int {
	for s.maxHeap.Len() > 0 && (*s.maxHeap)[0].value != s.prices[(*s.maxHeap)[0].time] {
		heap.Pop(s.maxHeap)
	}

	return (*s.maxHeap)[0].value
}

// Minimum and Maximum. Each pop takes log.N time so for N function calls this would take N.log.N time.
func (s *StockPrice) Minimum() int {
	for s.minHeap.Len() > 0 && (*s.minHeap)[0].value != s.prices[(*s.minHeap)[0].time] {
		heap.Pop(s.minHeap)
	}

	return (*s.minHeap)[0].value
}

type Price struct {
	time  int
	value int
}

type MinHeap []Price

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Price))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []Price

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i].value > h[j].value }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Price))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
