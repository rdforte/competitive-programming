package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(minimumAverage([]int{7, 8, 3, 4, 15, 13, 4, 1}) == 5.5)
	fmt.Println(minimumAverage([]int{7, 8, 3, 4, 15, 13, 4, 1}) == 5.5)
	fmt.Println(minimumAverage([]int{1, 2, 3, 7, 8, 9}) == 5.0)
}

func minimumAverage(nums []int) float64 {
	slices.Sort(nums)

	minAvg := math.MaxFloat64

	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		minAvg = min(float64((nums[i]+nums[j]))/2, minAvg)
	}

	return minAvg
}
