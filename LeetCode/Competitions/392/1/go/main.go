package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2, 3, 4}))
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2, 1}))
	fmt.Println(longestMonotonicSubarray([]int{3, 2, 1}))
	fmt.Println(longestMonotonicSubarray([]int{1, 10, 10}))
}

func longestMonotonicSubarray(nums []int) int {
	longest := 1
	shortest := 1
	longestPrev := 1
	shortestPrev := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			longestPrev++
		} else {
			longestPrev = 1
		}
		if nums[i] < nums[i-1] {
			shortestPrev++
		} else {
			shortestPrev = 1
		}
		longest = int(math.Max(float64(longest), float64(longestPrev)))
		shortest = int(math.Max(float64(shortest), float64(shortestPrev)))
	}

	if longest > shortest {
		return longest
	}
	return shortest
}
