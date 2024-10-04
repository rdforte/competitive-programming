package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func minPathSum(grid [][]int) int {
	rowLen, colLen := len(grid), len(grid[0])
	path := make([][]int, 0, rowLen)
	for range rowLen {
		path = append(path, make([]int, colLen))
	}

	path[0][0] = grid[0][0]
	for i := 1; i < colLen; i++ {
		path[0][i] = grid[0][i] + path[0][i-1]
	}
	for i := 1; i < rowLen; i++ {
		path[i][0] = grid[i][0] + path[i-1][0]
	}

	for r := 1; r < rowLen; r++ {
		for c := 1; c < colLen; c++ {
			shortC := grid[r][c] + path[r][c-1]
			shortR := grid[r][c] + path[r-1][c]

			path[r][c] = min(shortR, shortC)
		}
	}

	return path[rowLen-1][colLen-1]
}
