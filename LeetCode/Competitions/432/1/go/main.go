package main

import "fmt"

func main() {
	fmt.Println(zigzagTraversal([][]int{
		{2, 1},
		{2, 1},
		{2, 1},
	}))
	fmt.Println(zigzagTraversal([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(zigzagTraversal([][]int{
		{1, 2, 1, 3},
		{5, 15, 7, 3},
		{10, 4, 14, 12},
	}))
}

func zigzagTraversal(grid [][]int) []int {
	res := make([]int, 0, len(grid))
	for r := 0; r < len(grid); r++ {
		isEven := r&1 == 0
		if isEven {
			for c := 0; c < len(grid[r]); c++ {
				if c&1 == 0 {
					res = append(res, grid[r][c])
				}
			}
		} else {
			for c := len(grid[r]) - 1; c >= 0; c-- {
				if c&1 == 1 {
					res = append(res, grid[r][c])
				}
			}
		}
	}

	return res
}
