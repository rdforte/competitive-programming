package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0

	for left < right {
		area := math.Min(float64(height[left]), float64(height[right])) * float64(right-left)
		maxArea = int(math.Max(float64(maxArea), area))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
