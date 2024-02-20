package main

import "fmt"

func main() {

	graph := [][]byte{
		{'#', '#', '#', '.', '#', 'E', '#'},
		{'#', '.', '.', '.', '.', '.', '#'},
		{'#', 'S', '#', '.', '#', '#', '#'},
	}

	seen := make([][]bool, 0, len(graph))
	for i := range graph {
		seen = append(seen, make([]bool, len(graph[i])))
	}

	var path [][]int

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var findPath func(row, col int) bool
	findPath = func(row, col int) bool {
		if row < 0 || row >= len(graph) || col < 0 || col >= len(graph[row]) || seen[row][col] || graph[row][col] == '#' {
			return false
		}

		path = append(path, []int{row, col})

		seen[row][col] = true

		if graph[row][col] == 'E' {
			return true
		}

		for _, dir := range directions {
			nextMove := []int{row + dir[0], col + dir[1]}
			if findPath(nextMove[0], nextMove[1]) {
				return true
			}
		}

		path = path[:len(path)-1]

		return false
	}

	findPath(2, 1)

	fmt.Println(path)

}
