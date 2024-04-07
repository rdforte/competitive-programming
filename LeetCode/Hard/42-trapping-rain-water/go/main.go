package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	maxL := height[0]
	maxR := height[len(height)-1]

	trapped := 0

	for l, r := 0, len(height)-1; l < r; {
		if maxL < maxR {
			l++
			maxL = int(math.Max(float64(maxL), float64(height[l])))
			t := int(math.Min(float64(maxL), float64(maxR))) - height[l]
			if t > 0 {
				trapped += t
			}
		} else {
			r--
			maxR = int(math.Max(float64(maxR), float64(height[r])))
			t := int(math.Min(float64(maxL), float64(maxR))) - height[r]
			if t > 0 {
				trapped += t
			}
		}
	}

	return trapped
}
