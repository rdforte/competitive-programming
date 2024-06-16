package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(houseRobberTopDown([]int{1, 2, 3, 1}) == houseRobberBottomUp([]int{1, 2, 3, 1}))
	fmt.Println(houseRobberTopDown([]int{2, 7, 9, 3, 1}) == houseRobberBottomUp([]int{2, 7, 9, 3, 1}))
}

// DP - Top down approach using memoization.
func houseRobberTopDown(nums []int) int {
	memo := make(map[int]int)

	var dp func(i int) int
	dp = func(i int) int {
		if i == 0 {
			return nums[0]
		}
		if i == 1 {
			return int(math.Max(float64(nums[1]), float64(nums[0])))
		}

		if _, ok := memo[i]; !ok {
			memo[i] = int(math.Max(float64(nums[i]+dp(i-2)), float64(dp(i-1))))
		}

		return memo[i]
	}

	return dp(len(nums) - 1)
}

func houseRobberBottomUp(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]float64, len(nums))
	dp[0] = float64(nums[0])
	dp[1] = math.Max(float64(nums[1]), float64(nums[0]))

	for i := 2; i < len(nums); i++ {
		dp[i] = math.Max(float64(nums[i])+dp[i-2], dp[i-1])
	}

	return int(dp[len(dp)-1])
}
