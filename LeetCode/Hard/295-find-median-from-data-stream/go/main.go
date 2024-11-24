package main

import (
	"container/heap"
	"fmt"
)

func main() {
	f := Constructor()
	// f.AddNum(1)
	// f.AddNum(2)
	// fmt.Println(f.FindMedian())
	// f.AddNum(3)
	// fmt.Println(f.FindMedian())
	f.AddNum(6)
	f.AddNum(10)
	f.AddNum(2)
	f.AddNum(6)
	// 2, 6, 6, 10
	fmt.Println(f.FindMedian())
}

type MedianFinder struct {
	lo *MaxHeap
	hi *MinHeap
}

func Constructor() MedianFinder {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	return MedianFinder{lo: maxHeap, hi: minHeap}
}

func (this *MedianFinder) AddNum(num int) {
	n := float64(num)
	heap.Push(this.lo, n)                 // put the max num at the top
	heap.Push(this.hi, heap.Pop(this.lo)) // move max num over to hi
	if this.hi.Len() > this.lo.Len() {
		heap.Push(this.lo, heap.Pop(this.hi))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return this.lo.Top().(float64)
	}
	return (this.lo.Top().(float64) + this.hi.Top().(float64)) / 2.0
}

type MinHeap []float64

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() any {
	return (*h)[0]
}

type MaxHeap []float64

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() any {
	return (*h)[0]
}
