package main

import "fmt"

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}) == 3)
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}) == 10)
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}) == -2)
}

func maxSubarraySumCircular(nums []int) int {
	totalNumsSum, arraySumMax, arraySumMin, curMax, curMin := nums[0], nums[0], nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		totalNumsSum += nums[i]

		curMax = max(curMax+nums[i], nums[i])
		arraySumMax = max(arraySumMax, curMax)

		curMin = min(curMin+nums[i], nums[i])
		arraySumMin = min(arraySumMin, curMin)
	}

	// arraySumMin has used all elements in array to create the min so not a viable solution.
	// sub array needs to have atleast 1 element in it.
	if totalNumsSum-arraySumMin == 0 {
		return arraySumMax
	}

	return max(arraySumMax, totalNumsSum-arraySumMin)
}
