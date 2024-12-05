package main

import "fmt"

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	ptrs := [][]int{
		{0, 0},                                // top left
		{0, len(matrix[0]) - 1},               // top right
		{len(matrix) - 1, len(matrix[0]) - 1}, // bottom right
		{len(matrix) - 1, 0},                  // bottom left
	}
	tl, tr, br, bl := ptrs[0], ptrs[1], ptrs[2], ptrs[3]

	for tl[0] < bl[0] {
		// rotate
		moveBy := (tr[1] - tl[1])
		for i := 0; i < moveBy; i++ {
			trTemp := matrix[tr[0]+i][tr[1]]
			matrix[tr[0]+i][tr[1]] = matrix[tl[0]][tl[1]+i]

			brTemp := matrix[br[0]][br[1]-i]
			matrix[br[0]][br[1]-i] = trTemp

			blTemp := matrix[bl[0]-i][bl[1]]
			matrix[bl[0]-i][bl[1]] = brTemp

			matrix[tl[0]][tl[1]+i] = blTemp
		}

		// move pointers in
		tl[0], tl[1] = tl[0]+1, tl[1]+1
		tr[0], tr[1] = tr[0]+1, tr[1]-1
		br[0], br[1] = br[0]-1, br[1]-1
		bl[0], bl[1] = bl[0]-1, bl[1]+1
	}
}
