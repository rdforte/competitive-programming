package main

import (
	"fmt"
	"math"
)

const (
	EXIT  = 1
	ENTER = 0
)

func main() {
	// arrival, action
	fmt.Println(timeTake([]int{0, 1, 3, 5}, []int{1, 1, 1, 1}))
	fmt.Println(timeTake([]int{0, 1, 1, 2, 4}, []int{0, 1, 0, 0, 1}))
	fmt.Println(timeTake([]int{0, 0, 0}, []int{1, 0, 1}))
	fmt.Println(timeTake([]int{3, 3, 4, 5, 5, 5}, []int{1, 0, 1, 0, 1, 0})) // UNABLE TO SOLVE FOR THIS CASE
}

func timeTake(arrival, state []int) []int {
	clock := 0

	res := make([]int, len(state))

	processed := 0
	prevAction := -1 // -1 is door not used, 0 enter, 1 exit

	for processed < len(state) {

		persons := make([]int, 0, 1) // index
		for i, a := range arrival {
			if a == clock {
				persons = append(persons, i)
			}
		}

		if len(persons) == 0 {
			prevAction = -1
		}

		if len(persons) == 1 {
			i := persons[0]
			res[i] = clock
			processed++
			prevAction = state[i]
		}

		if len(persons) > 1 {
			nextPerson := math.MaxInt
			for _, p := range persons {
				if (prevAction == -1 || prevAction == EXIT) && state[p] == EXIT {
					if p < nextPerson {
						nextPerson = p
					} else {
						arrival[p] = clock + 1
					}
				} else if prevAction == ENTER && state[p] == ENTER {
					if p < nextPerson {
						nextPerson = p
					} else {
						arrival[p] = clock + 1
					}
				} else {
					arrival[p] = clock + 1
				}
			}
			res[nextPerson] = clock
			processed++
			prevAction = state[nextPerson]
		}

		clock++
	}

	return res
}

// priority queue with index and arrival time

type MinHeap [][]int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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
