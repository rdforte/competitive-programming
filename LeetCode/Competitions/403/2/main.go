package main

import "fmt"

func main() {
	fmt.Println(minimumArea([][]int{
		{0, 1, 0},
		{1, 0, 1},
	}))
	fmt.Println(minimumArea([][]int{
		{0, 0},
		{1, 0},
	}))
}

func minimumArea(grid [][]int) int {
	top, right, bottom, left := len(grid)-1, 0, 0, len(grid[0])-1

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				if r < top {
					top = r
				}
				if r > bottom {
					bottom = r
				}
				if c < left {
					left = c
				}
				if c > right {
					right = c
				}
			}
		}
	}

	height := (bottom - top) + 1
	width := (right - left) + 1

	return height * width
}
