package main

import (
	"math"
	"sort"
)

func main() {
	println(minimumDiff([]int{9, 4, 1, 7}, 2))
}

// cur
//
// 1, 4, 7, 9

func minimumDiff(nums []int, k int) int {
	sort.Ints(nums)
	minRes := math.MaxInt

	for l, r := 0, k-1; r < len(nums); l, r = l+1, r+1 {
		minRes = min(minRes, nums[r]-nums[l])
	}

	return minRes
}
