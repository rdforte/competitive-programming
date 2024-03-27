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
	t, r, b, l := 0, len(matrix[0])-1, len(matrix)-1, 0

	for l < r {
		for i := 0; i < r-l; i++ {
			tr := matrix[t+i][r]
			matrix[t+i][r] = matrix[t][l+i]

			br := matrix[b][r-i]
			matrix[b][r-i] = tr

			bl := matrix[b-i][l]
			matrix[b-i][l] = br

			matrix[t][l+i] = bl
		}
		t++
		b--
		l++
		r--
	}
}
