package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
}

func largestRectangleArea(heights []int) int {
	stack := [][]int{
		{0, heights[0]},
	}

	maxHeight := heights[0]
	for i := 1; i < len(heights); i++ {
		curHeight := heights[i]

		barPos := i
		for len(stack) > 0 && curHeight < stack[len(stack)-1][1] {
			bar := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i - bar[0]
			area := width * bar[1]
			maxHeight = int(math.Max(float64(maxHeight), float64(area)))
			barPos = bar[0]
		}
		stack = append(stack, []int{barPos, heights[i]})
	}

	for len(stack) > 0 {
		bar := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := ((len(heights) - 1) - bar[0]) + 1
		area := width * bar[1]
		maxHeight = int(math.Max(float64(maxHeight), float64(area)))
	}

	return maxHeight
}
