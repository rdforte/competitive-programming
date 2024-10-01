package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	grid := make([][]int, 0, m)
	for range m {
		grid = append(grid, make([]int, n))
	}
	grid[0][0] = 1

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r > 0 {
				grid[r][c] += grid[r-1][c]
			}
			if c > 0 {
				grid[r][c] += grid[r][c-1]
			}
		}
	}

	return grid[m-1][n-1]
}
