package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1},
		{0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rowLen := len(obstacleGrid)
	colLen := len(obstacleGrid[0])

	paths := make([][]int, 0, rowLen)
	for range rowLen {
		paths = append(paths, make([]int, colLen))
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	paths[0][0] = 1

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if obstacleGrid[r][c] == 1 {
				continue
			}
			if r > 0 {
				paths[r][c] += paths[r-1][c]
			}
			if c > 0 {
				paths[r][c] += paths[r][c-1]
			}
		}
	}

	return paths[rowLen-1][colLen-1]
}
