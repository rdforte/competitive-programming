package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{8, -19, 5, -4, 20}))
}

func maxSubArray(nums []int) int {
	curSubArray, maxSubArray := nums[0], nums[0]

	for _, num := range nums[1:] {
		// If current_subarray is negative, throw it away. Otherwise, keep
		// adding to it.
		curSubArray = max(num+curSubArray, num)
		maxSubArray = max(curSubArray, maxSubArray)
	}

	return maxSubArray
}
