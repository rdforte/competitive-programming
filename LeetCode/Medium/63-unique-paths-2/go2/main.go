package main

import (
	"fmt"
)

func main() {
	// println(uniquePathsWithObstacles([][]int{
	// {0, 0, 0},
	// {0, 1, 0},
	// {0, 0, 0},
	// }))

	println(uniquePathsWithObstacles([][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}))

	println(uniquePathsWithObstaclesNoMemo([][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	numRows, numCols := len(obstacleGrid), len(obstacleGrid[0])

	memo := make([][]int, numRows)
	for r := range numRows {
		memo[r] = make([]int, numCols)
		for c := range numCols {
			memo[r][c] = -1
		}
	}

	var moveRecursive func(row, col int) int
	moveRecursive = func(row, col int) int {
		if row >= numRows || col >= numCols || obstacleGrid[row][col] == 1 {
			return 0
		}

		if row == numRows-1 && col == numCols-1 {
			return 1
		}

		if memo[row][col] == -1 {
			memo[row][col] = moveRecursive(row+1, col) + moveRecursive(row, col+1)
		}

		return memo[row][col]
	}

	return moveRecursive(0, 0)
}

func uniquePathsWithObstaclesNoMemo(obstacleGrid [][]int) int {
	numRows, numCols := len(obstacleGrid), len(obstacleGrid[0])

	iterations := 0

	var moveRecursive func(row, col int) int
	moveRecursive = func(row, col int) int {
		iterations++
		if row >= numRows || col >= numCols || obstacleGrid[row][col] == 1 {
			return 0
		}

		if row == numRows-1 && col == numCols-1 {
			return 1
		}

		return moveRecursive(row+1, col) + moveRecursive(row, col+1)

	}

	res := moveRecursive(0, 0)
	fmt.Printf("iters: %d\n", iterations)
	return res
}
