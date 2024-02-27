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

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for key, val := range mp {
		heap.Push(h, []int{val, key})
	}

	res := []int{}

	for i := 0; i < k && h.Len() > 0; i++ {
		el := heap.Pop(h).([]int)
		res = append(res, el[1])
	}

	return res
}

type MinHeap [][]int

func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Less(i, j int) bool { return (*h)[i][0] > (*h)[j][0] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
