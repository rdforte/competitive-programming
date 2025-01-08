package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}

// NOTE: Refer to prefix sum. Prefix sum is a good use case for this question because
// we can get the total number of contiguous sub arrays.
// This is a combination of Prefix sum, sliding window and some clever thinking.

// Time O(n)
// Space O(1)
func numSubarraysWithSum(nums []int, goal int) int {
	return countNumSubArrays(nums, goal) - countNumSubArrays(nums, goal-1)
}

func countNumSubArrays(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}

	res, sum := 0, 0

	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]

		for sum > goal {
			sum -= nums[l]
			l++
		}

		res += ((r - l) + 1)
	}

	return res
}
