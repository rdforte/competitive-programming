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

	fmt.Println("TOPOLOGICAL SORT")
	fmt.Println(canFinishTopologicalSort(2, [][]int{
		{1, 0},
	}))

	fmt.Println(canFinishTopologicalSort(2, [][]int{
		{1, 0},
		{0, 1},
	}))

	fmt.Println(canFinishTopologicalSort(3, [][]int{
		{1, 0},
		{0, 2},
		{2, 1},
	}))
	fmt.Println(canFinishTopologicalSort(3, [][]int{
		{1, 0},
		{1, 2},
		{0, 1},
	}))
}

func canFinishTopologicalSort(numCourses int, prerequisites [][]int) bool {
	// initialise adjacency list takes O(m) time where m is the edges
	indegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		indegree[p[0]]++
		adj[p[1]] = append(adj[p[1]], p[0])
	}

	var queue []int
	for i, in := range indegree {
		if in == 0 {
			queue = append(queue, i)
		}
	}

	visited := 0

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		visited++

		for _, edge := range adj[n] {
			indegree[edge]--
			if indegree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}

	return visited == numCourses
}

// time and space = O(m+n)

func canFinish(numCourses int, prerequisites [][]int) bool {
	// initialise adjacency list takes O(m) time where m is the edges
	courses := make([][]int, numCourses)
	for _, p := range prerequisites {
		courses[p[0]] = append(courses[p[0]], p[1])
	}

	visitedBefore := make([]bool, numCourses)

	// we iterate through each node once O(n) time
	for i := range numCourses {
		seen := make([]bool, numCourses)

		var isPossible func(n int) bool
		isPossible = func(n int) bool {
			if seen[n] {
				return false
			}

			// We need this because it saves us from having to traverse the same path over again
			// which will save time.
			if visitedBefore[n] {
				return true
			}

			seen[n] = true
			visitedBefore[n] = true

			for _, pre := range courses[n] {
				if !isPossible(pre) {
					return false
				}
			}
			seen[n] = false

			return true
		}

		if !isPossible(i) {
			return false
		}

	}

	return true
}
