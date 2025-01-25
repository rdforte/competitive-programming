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

	fmt.Println(exist([][]byte{
		{'a', 'a', 'b', 'a', 'a', 'b'},
		{'a', 'a', 'b', 'b', 'b', 'a'},
		{'a', 'a', 'a', 'a', 'b', 'a'},
		{'b', 'a', 'b', 'b', 'a', 'b'},
		{'a', 'b', 'b', 'a', 'b', 'a'},
		{'b', 'a', 'a', 'a', 'a', 'b'},
	}, "bbbaabbbbbab") == false)
}

// Instead of trying to match on the full word, this can lead to a very INEFFICIENT algorithm so
// instead we can match one letter at a time as we progress throughout the grid. If the letter for that
// index does not match then we know that this chain of letters can not be a possible word so no need to
// progress any further and instead return early.

func exist(board [][]byte, word string) bool {
	dirs := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			var dp func(r, c, i int) bool
			dp = func(r, c, i int) bool {
				if i == len(word) {
					return true
				}

				if r >= len(board) || r < 0 || c < 0 || c >= len(board[r]) || board[r][c] == '.' {
					return false
				}

				if board[r][c] != word[i] {
					return false
				}

				char := board[r][c]
				board[r][c] = '.'

				for _, d := range dirs {
					if dp(r+d[0], c+d[1], i+1) {
						return true
					}
				}

				board[r][c] = char

				return false
			}

			if dp(r, c, 0) {
				return true
			}
		}
	}

	return false
}
