package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
}

func islandPerimeter(grid [][]int) int {

	perim := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				// look left
				if c == 0 || grid[r][c-1] == 0 {
					perim++
				}

				// look up
				if r == 0 || grid[r-1][c] == 0 {
					perim++
				}

				// look right
				if c == len(grid[r])-1 || grid[r][c+1] == 0 {
					perim++
				}

				// look down
				if r == len(grid)-1 || grid[r+1][c] == 0 {
					perim++
				}
			}
		}
	}

	return perim
}
