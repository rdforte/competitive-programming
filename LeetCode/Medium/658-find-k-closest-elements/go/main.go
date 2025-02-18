package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElements([]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2))
	fmt.Println("-----------------")
	fmt.Println(findClosestElementsV2([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElementsV2([]int{1, 1, 2, 3, 4, 5}, 4, -1))
	fmt.Println(findClosestElementsV2([]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8}, 2, 2))
}

func findClosestElementsV2(arr []int, k, x int) []int {
	if len(arr) <= k {
		return arr
	}

	left, right, sum := 0, 0, math.MaxInt

	curSum := 0
	for l, r := 0, 0; r < len(arr); r++ {
		curSum += int(math.Abs(float64(arr[r]) - float64(x)))
		if r-l >= k {
			curSum -= int(math.Abs(float64(arr[l]) - float64(x)))
			l++
		}

		if r-l == k-1 && curSum < sum {
			sum = curSum
			left = l
			right = r
		}
	}

	return arr[left : right+1]
}

func findClosestElements(arr []int, k, x int) []int {
	h := &IntHeap{x: x, arr: make([]int, 0, len(arr))}
	heap.Init(h)

	for _, a := range arr {
		heap.Push(h, a)
	}

	res := make([]int, 0, k)
	for i := 0; i < k && h.Len() > 0; i++ {
		res = append(res, heap.Pop(h).(int))
	}

	sort.Ints(res)
	return res
}

type IntHeap struct {
	x   int
	arr []int
}

func (h IntHeap) Len() int { return len(h.arr) }

func (h IntHeap) Less(i, j int) bool {
	if math.Abs(float64(h.x)-float64(h.arr[i])) == math.Abs(float64(h.x)-float64(h.arr[j])) {
		return h.arr[i] < h.arr[j]
	}
	return math.Abs(float64(h.x)-float64(h.arr[i])) < math.Abs(float64(h.x)-float64(h.arr[j]))
}

func (h IntHeap) Swap(i, j int) { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

func (h *IntHeap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}
