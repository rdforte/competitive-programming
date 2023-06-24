package validSudoku

func IsValidSudoku(board [][]byte) bool {
	type void struct{}
	col := make(map[int]map[string]void)
	row := make(map[int]map[string]void)
	box := make(map[int]map[int]map[string]void)

	for r := 0; r < len(board); r++ {
		row[r] = make(map[string]void)
		for c := 0; c < len(board[r]); c++ {
			if r == 0 {
				col[c] = make(map[string]void)
			}

			num := string(board[r][c])
			if num == "." {
				continue
			}

			if _, ok := row[r][num]; ok {
				return false
			} else {
				row[r][num] = void{}
			}

			if _, ok := col[c][num]; ok {
				return false
			} else {
				col[c][num] = void{}
			}

			boxRow := (r / 3) % 3
			boxCol := (c / 3) % 3

			if val := box[boxRow]; val == nil {
				box[boxRow] = make(map[int]map[string]void)
			}

			if val := box[boxRow][boxCol]; val == nil {
				box[boxRow][boxCol] = make(map[string]void)
			}

			if _, ok := box[boxRow][boxCol][num]; ok {
				return false
			} else {
				box[boxRow][boxCol][num] = void{}
			}
		}
	}

	return true
}
