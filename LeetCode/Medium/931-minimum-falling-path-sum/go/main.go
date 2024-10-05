package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
	fmt.Println(minFallingPathSum([][]int{
		{-19, 57},
		{-40, 5},
	}))
}

func minFallingPathSum(matrix [][]int) int {
	rowLen, colLen := len(matrix), len(matrix[0])
	for r := 1; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			above := matrix[r-1][c] + matrix[r][c]
			aboveLeft := math.MaxInt
			aboveRight := math.MaxInt
			if c > 0 {
				aboveLeft = matrix[r][c] + matrix[r-1][c-1]
			}
			if c < colLen-1 {
				aboveRight = matrix[r][c] + matrix[r-1][c+1]
			}
			matrix[r][c] = min(above, aboveLeft, aboveRight)
		}
	}

	minFallingPath := matrix[rowLen-1][colLen-1]
	for _, n := range matrix[rowLen-1] {
		minFallingPath = min(minFallingPath, n)
	}

	return minFallingPath
}
