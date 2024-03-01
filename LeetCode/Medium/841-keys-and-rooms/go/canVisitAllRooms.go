package visitRooms

func CanVisitAllRooms(rooms [][]int) bool {
	stack := []int{}
	stack = append(stack, rooms[0]...)
	rooms[0] = []int{}
	visited := make([]bool, len(rooms))
	visited[0] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited[node] = true

		stack = append(stack, rooms[node]...)
		rooms[node] = []int{}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
