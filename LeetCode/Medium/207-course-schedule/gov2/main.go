package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{
		{1, 0},
	}))

	fmt.Println(canFinish(2, [][]int{
		{1, 0},
		{0, 1},
	}))

	fmt.Println(canFinish(3, [][]int{
		{1, 0},
		{0, 2},
		{2, 1},
	}))
	fmt.Println(canFinish(3, [][]int{
		{1, 0},
		{1, 2},
		{0, 1},
	}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int, numCourses)
	indirection := make([]int, numCourses)
	for _, req := range prerequisites {
		adj[req[1]] = append(adj[req[1]], req[0])
		indirection[req[0]]++
	}

	var queue []int
	for course, numIndirection := range indirection {
		if numIndirection == 0 {
			queue = append(queue, course)
		}
	}

	completedCourses := 0
	for len(queue) > 0 {
		completedCourses++
		course := queue[0]
		queue = queue[1:]

		for _, c := range adj[course] {
			indirection[c]--
			if indirection[c] == 0 {
				queue = append(queue, c)
			}
		}
	}

	return completedCourses == numCourses
}
