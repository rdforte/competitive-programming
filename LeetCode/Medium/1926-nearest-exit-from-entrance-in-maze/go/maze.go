package nearestExit

func NearestExit(maze [][]byte, entrance []int) int {
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	seen := make([][]bool, len(maze))
	steps := make([][]int, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[i]))
		steps[i] = make([]int, len(maze[i]))
	}

	queue := [][]int{}

	queue = append(queue, entrance)

	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		if seen[row][col] || maze[row][col] == '+' {
			continue
		}

		seen[row][col] = true

		isEntrance := row == entrance[0] && col == entrance[1]
		isEdgeOfMaze := row == 0 || row == len(maze)-1 || col == 0 || col == len(maze[row])-1

		if isEdgeOfMaze && !isEntrance {
			return steps[row][col]
		}

		for _, dir := range directions {
			r := row + dir[0]
			c := col + dir[1]

			if r >= 0 && r < len(maze) && c >= 0 && c < len(maze[r]) {
				steps[r][c] = steps[row][col] + 1
				queue = append(queue, []int{r, c})
			}
		}

	}

	return -1
}
