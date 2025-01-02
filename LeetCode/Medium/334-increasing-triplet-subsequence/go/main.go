package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}) == true)
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}) == false)
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}) == true)
}

func increasingTriplet(nums []int) bool {
	i, j := math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n > j {
			return true
		}

		if n < i {
			i = n
			continue
		}

		if n > i && n < j {
			j = n
		}
	}

	return false
}
