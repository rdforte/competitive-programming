package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))
	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	res := 1
	runningSum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		runningSum += nums[r]
		maxSum := ((r - l) + 1) * nums[r]

		for maxSum-runningSum > k {
			runningSum -= nums[l]
			l++
			maxSum = ((r - l) + 1) * nums[r]
		}

		res = max(res, (r-l)+1)
	}

	return res
}
