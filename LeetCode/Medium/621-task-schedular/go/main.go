package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	fmt.Println(leastInterval([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
}

func leastInterval(tasks []byte, n int) int {
	h := &IntHeap{}
	heap.Init(h)

	tally := [26]int{}
	for _, v := range tasks {
		tally[v-'A']++
	}
	for _, v := range tally {
		if v > 0 {
			heap.Push(h, v)
		}
	}

	queue := [][]int{}
	time := 0

	for h.Len() > 0 || len(queue) > 0 {
		time++
		if h.Len() > 0 {
			top := heap.Pop(h).(int)
			if top > 1 {
				queue = append(queue, []int{top - 1, time + n})
			}
		}

		if len(queue) > 0 && queue[0][1] == time {
			heap.Push(h, queue[0][0])
			queue = queue[1:]
		}
	}

	return time
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
