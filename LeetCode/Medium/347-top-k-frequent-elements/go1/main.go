package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
}

// Time complexity =
// complexity of min heap log.N
// elements to retrieve = k
// therefore N.log.K
func topKFrequent(nums []int, k int) []int {
	uniqueNums := make(map[int]int, len(nums))
	for _, n := range nums {
		uniqueNums[n]++
	}

	h := &IntHeap{}
	heap.Init(h)
	for k, v := range uniqueNums {
		heap.Push(h, []int{k, v})
	}

	res := make([]int, 0, k)
	for i := 0; i < k && h.Len() > 0; i++ {
		res = append(res, heap.Pop(h).([]int)[0])
	}

	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }

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
