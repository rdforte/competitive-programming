package main

import "fmt"

func main() {
	fmt.Println(minDays([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}))
	fmt.Println(minDays([][]int{
		{0, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
		{1, 1, 1, 1},
	}))
}

func minDays(grid [][]int) int {
	islandMap := make([][]int, 0, len(grid))
	for i := range len(grid) {
		islandMap = append(islandMap, make([]int, len(grid[i])))
	}

	directions := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	var numIslands int
	var waters [][]int

	var findIslandDFS func(r, c, island int)
	findIslandDFS = func(r, c, island int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
			return // out of bounds
		}

		if grid[r][c] == 0 {
			waters = append(waters, []int{r, c}) // keep track of water
			return                               // water
		}

		if islandMap[r][c] != 0 {
			return // already found land before and marked it
		}

		// have not seen this land before. mark it
		islandMap[r][c] = island

		// find neighbouring lands
		for _, dir := range directions {
			findIslandDFS(r+dir[0], c+dir[1], island)
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				continue // water
			}

			if islandMap[r][c] != 0 {
				continue // already seen island
			}

			// have not seen this piece of land before, find out which island it belongs to.
			numIslands++
			findIslandDFS(r, c, numIslands)
		}
	}

	// for _, i := range islandMap {
	// fmt.Println(i)
	// }

	if numIslands >= 2 || numIslands == 0 {
		return 0
	}

	// check to see if can change single land to water.
	for _, w := range waters {
	}

	return 0
}
