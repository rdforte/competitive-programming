package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED") == true)

	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE") == true)

	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB") == false)

	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCESEEEFS") == true)

	fmt.Println(exist([][]byte{
		{'a', 'a'},
	}, "aaa") == false)
}

func exist(board [][]byte, word string) bool {

	seen := false

	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

LOOP:
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			visited := [][]bool{}
			for i := 0; i < len(board); i++ {
				visited = append(visited, make([]bool, len(board[0])))
			}

			visited[row][col] = true

			var dfs func(r, c, i int)
			dfs = func(r, c, i int) {
				if seen {
					return
				}

				if word[i] != board[r][c] {
					return
				}

				if i == len(word)-1 {
					seen = true
					return
				}

				for _, dir := range directions {
					nextR := r + dir[0]
					nextC := c + dir[1]
					if nextR >= len(board) || nextR < 0 || nextC < 0 || nextC >= len(board[r]) {
						continue
					}

					if visited[nextR][nextC] {
						continue
					}

					visited[nextR][nextC] = true
					dfs(nextR, nextC, i+1)
					visited[nextR][nextC] = false
				}

			}

			dfs(row, col, 0)

			if seen {
				break LOOP
			}
		}
	}

	return seen
}
