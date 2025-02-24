package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1, 10, 7, 1, 1, 4, 2, 2, 3}, 4))
}

// n*log*k + n + k
// n*log*k
func topKFrequent(nums []int, k int) []int {
	// n = len nums
	// n time
	elem := make(map[int]int)
	for _, n := range nums {
		elem[n]++
	}

	h := &IntHeap{}
	heap.Init(h)

	// n * log * k
	for key, val := range elem {
		heap.Push(h, []int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// k
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).([]int)[0]
	}

	return res
}

type IntHeap [][]int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
