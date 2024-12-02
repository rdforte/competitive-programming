package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getLargestOutlier([]int{2, 3, 5, 10}) == 10)
	fmt.Println(getLargestOutlier([]int{-2, -1, -3, -6, 4}) == 4)
	fmt.Println(getLargestOutlier([]int{-310, -702, -702}) == -310)
	fmt.Printf("\nEDGE CASE 1 --------------------------\n\n")
	// edge case. Half doubleSum = outlier
	fmt.Println(getLargestOutlier([]int{6, -31, 50, -35, 41, 37, -42, 13}) == -35) // sum: 37, outlier -35
	fmt.Printf("\nEDGE CASE 2 --------------------------\n\n")
	// if we take into account edge case of half the doubleSum = outlier then we might run into the issue
	// of the outlier being the same value as the sum in which case we need the frequency.
	fmt.Println(getLargestOutlier([]int{1, 1, 1, 1, 1, 5, 5}) == 5)
}

func getLargestOutlier(nums []int) int {
	sum := 0
	seen := make(map[int]int, len(nums))
	for _, n := range nums {
		seen[n]++
		sum += n
	}

	res := math.MinInt

	for _, outlier := range nums {
		doubleSum := sum - outlier
		if doubleSum%2 != 0 {
			continue
		}

		num := doubleSum / 2
		if freq, ok := seen[num]; (ok && num != outlier) || (ok && freq > 1) {
			res = max(res, outlier)
		}
	}

	return res
}
