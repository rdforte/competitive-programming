package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMaxSum([]int{18, 11, 24, 9, 10, 11, 7, 29, 16}, []int{28, 26, 27, 4, 2, 19, 23, 1, 2}, 1))
	fmt.Println(findMaxSum([]int{4, 2, 1, 5, 3}, []int{10, 20, 30, 40, 50}, 2))
	fmt.Println(findMaxSum([]int{2, 2, 2, 2}, []int{3, 1, 2, 3}, 1))
	// fmt.Println(findMaxSum([]int{2, 2, 2, 2, 3}, []int{3, 1, 2, 3, 4}, 1))
	// fmt.Println(findMaxSum([]int{2, 2, 2, 2, 1}, []int{3, 1, 2, 3, 4}, 1))
}

func findMaxSum(nums1 []int, nums2 []int, k int) []int64 {
	pairs := make([][]int, 0, len(nums1)) // idx, weight

	for i, n := range nums1 {
		pairs = append(pairs, []int{i, n})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	h := &IntHeap{}
	heap.Init(h)

	sum := 0
	res := make([]int64, len(nums1))
	for i := 1; i < len(pairs); i++ {
		idx := pairs[i-1][0]
		heap.Push(h, nums2[idx])
		sum += nums2[idx]

		if h.Len() > k {
			sum -= heap.Pop(h).(int)
		}

		// if same weight then use result of previous calculation.
		if pairs[i][1] == pairs[i-1][1] {
			res[pairs[i][0]] = res[pairs[i-1][0]]
			continue
		}

		res[pairs[i][0]] = int64(sum)
	}

	return res
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
