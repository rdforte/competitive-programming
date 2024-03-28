package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestConsecutive([]int{1, 3, 5, 2}))
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}

// Its faster just to sort here instead of using a minHeap

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}

	maxConsec := 1
	count := 1
	prev := heap.Pop(h)
	for h.Len() > 0 {
		top := heap.Pop(h).(int)
		if top == prev {
			continue
		}
		if top-1 == prev {
			count++
		} else {
			count = 1
		}
		prev = top
		maxConsec = int(math.Max(float64(maxConsec), float64(count)))
	}

	return maxConsec
}

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
