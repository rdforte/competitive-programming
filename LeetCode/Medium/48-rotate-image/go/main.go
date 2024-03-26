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
	rotateCorners(matrix)

	// rotate top to right side
	offset := 1

	queue := []int{}
	for i := offset; i < len(matrix)-offset; i++ {
		queue = append(queue, matrix[0][i])
	}

	for i := offset; i < len(matrix)-offset; i++ {
		queue = append(queue, matrix[i][len(matrix)-1])
	}

	for c := offset; c < len(matrix)-offset; c++ {
		matrix[c][len(matrix)-1] = queue[0]
		queue = queue[1:]
	}

	// move right side to bottom
	for i := len(matrix) - 1 - offset; i >= offset; i-- {
		queue = append(queue, matrix[len(matrix)-1][i])
	}

	for r := len(matrix) - 1 - offset; r >= offset; r-- {
		matrix[len(matrix)-1][r] = queue[0]
		queue = queue[1:]
	}

	fmt.Println(queue)
}

func rotateCorners(matrix [][]int) {
	topLeft := []int{0, 0}
	topRight := []int{0, len(matrix[0]) - 1}
	bottomRight := []int{len(matrix) - 1, len(matrix[len(matrix)-1]) - 1}
	bottomLeft := []int{len(matrix) - 1, 0}

	for topLeft[0] < bottomRight[0] {
		oldTopRight := matrix[topRight[0]][topRight[1]]
		matrix[topRight[0]][topRight[1]] = matrix[topLeft[0]][topLeft[1]]

		oldBottomRight := matrix[bottomRight[0]][bottomRight[1]]
		matrix[bottomRight[0]][bottomRight[1]] = oldTopRight

		oldBottomLeft := matrix[bottomLeft[0]][bottomLeft[1]]
		matrix[bottomLeft[0]][bottomLeft[1]] = oldBottomRight

		matrix[topLeft[0]][topLeft[1]] = oldBottomLeft

		topLeft[0]++
		topLeft[1]++

		topRight[0]++
		topRight[1]--

		bottomRight[0]--
		bottomRight[1]--

		bottomLeft[0]--
		bottomLeft[1]++
	}
}
