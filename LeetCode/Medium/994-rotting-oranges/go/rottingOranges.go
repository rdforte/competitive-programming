package rottingOranges

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func OrangesRotting(grid [][]int) int {

	queue := [][]int{}
	freshOrangeCount := 0
	minutes := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				freshOrangeCount++
			}
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	for len(queue) > 0 && freshOrangeCount > 0 {
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			orange := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				r := orange[0] + dir[0]
				c := orange[1] + dir[1]
				if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) && grid[r][c] == 1 {
					queue = append(queue, []int{r, c})
					grid[r][c] = 2
					freshOrangeCount--
				}
			}
		}
		minutes++
	}

	if freshOrangeCount == 0 {
		return minutes
	}

	return -1
}
