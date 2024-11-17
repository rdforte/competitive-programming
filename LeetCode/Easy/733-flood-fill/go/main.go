package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}, 1, 1, 2))
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	startingPixelColor := image[sr][sc]
	if startingPixelColor == color {
		return image // grid is already coloured and doesn't need changing.
	}

	var fill func(r, c int)
	fill = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[r]) || image[r][c] != startingPixelColor {
			return
		}

		image[r][c] = color

		for _, d := range directions {
			fill(r+d[0], c+d[1])
		}
	}

	fill(sr, sc)

	return image
}
