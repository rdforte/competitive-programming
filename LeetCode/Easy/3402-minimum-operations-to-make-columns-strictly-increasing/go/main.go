package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([][]int{
		{3, 2},
		{1, 3},
		{3, 4},
		{0, 1},
	}))
	fmt.Println(minimumOperations([][]int{
		{3, 2, 1},
		{2, 1, 0},
		{1, 2, 3},
	}))
}

func minimumOperations(grid [][]int) int {
	ops := 0
	for c := 0; c < len(grid[0]); c++ {
		for r := 1; r < len(grid); r++ {
			incrimentBy := max(0, (grid[r-1][c]-grid[r][c])+1)
			ops += incrimentBy
			grid[r][c] += incrimentBy
		}
	}
	return ops
}
