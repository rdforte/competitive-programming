package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 10))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 0))
	fmt.Println(topKFrequent([]int{}, 1))
}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}

	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for key, val := range mp {
		heap.Push(minHeap, []int{key, val})

		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	res := make([]int, minHeap.Len())
	for i := minHeap.Len() - 1; i >= 0; i-- {
		res[i] = heap.Pop(minHeap).([]int)[0]
	}

	return res
}

type IntHeap [][]int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
