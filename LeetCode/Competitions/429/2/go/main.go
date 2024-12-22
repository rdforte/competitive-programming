package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2))
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4}, 1))
	fmt.Println(maxDistinctElements([]int{4, 4, 4, 4, 3}, 0))
	fmt.Println(maxDistinctElements([]int{6, 7, 6, 6, 7}, 1))
}

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)

	res := 0
	prev := math.MinInt64

	for i := 0; i < len(nums); i++ {
		lowerBound := nums[i] - k
		upperBound := nums[i] + k

		if lowerBound > prev {
			res++
			prev = lowerBound
		} else if prev >= lowerBound && prev < upperBound {
			prev++
			res++
		}
	}

	return res
}
