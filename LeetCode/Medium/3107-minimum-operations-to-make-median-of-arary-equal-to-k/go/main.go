package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 4))
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 7))
	fmt.Println(minOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4))
}

// NOTE: this questions assumes the median is just len(nums)/2
/*
* For all elements before mid, we ensure that those are at most k.
For all elements after mid, we esure that those are at least k.
The element at mid is our "median", so we ensure that it is exactly k.
*/
func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)

	mid := len(nums) / 2
	var operations int

	for i := mid - 1; i >= 0 && nums[i] > k; i-- {
		operations += nums[i] - k
	}
	for i := mid + 1; i < len(nums) && nums[i] < k; i++ {
		operations += k - nums[i]
	}

	operations += int(math.Abs(float64(k - nums[mid])))

	return int64(operations)
}
