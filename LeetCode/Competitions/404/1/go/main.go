package main

import "fmt"

func main() {
	fmt.Println(maxHeightOfTriangle(2, 4) == 3)
	fmt.Println(maxHeightOfTriangle(2, 1) == 2)
	fmt.Println(maxHeightOfTriangle(1, 1) == 1)
	fmt.Println(maxHeightOfTriangle(10, 1) == 2)
	fmt.Println(maxHeightOfTriangle(4, 4) == 3)
	fmt.Println(maxHeightOfTriangle(10, 10))
	fmt.Println(maxHeightOfTriangle(4, 9))
}

func maxHeightOfTriangle(red int, blue int) int {
	var maxLevel func(r, b, l int) int
	maxLevel = func(r, b, l int) int {
		if b-l < 0 {
			return l - 1
		}

		sr, sb := b, r
		return maxLevel(sr-l, sb, l+1)
	}

	return max(maxLevel(red, blue, 1), maxLevel(blue, red, 1))
}
