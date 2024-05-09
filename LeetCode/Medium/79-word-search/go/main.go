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
			var dfs func(r, c, i int)
			dfs = func(r, c, i int) {
				if seen {
					return
				}

				if r >= len(board) || r < 0 || c < 0 || c >= len(board[r]) {
					return
				}

				if board[r][c] == '#' {
					return
				}

				if word[i] != board[r][c] {
					return
				}

				if i == len(word)-1 {
					seen = true
					return
				}

				cur := board[r][c]
				board[r][c] = '#'

				for _, dir := range directions {
					nextR := r + dir[0]
					nextC := c + dir[1]

					dfs(nextR, nextC, i+1)
				}

				board[r][c] = cur

			}

			dfs(row, col, 0)

			if seen {
				break LOOP
			}
		}
	}

	return seen
}
