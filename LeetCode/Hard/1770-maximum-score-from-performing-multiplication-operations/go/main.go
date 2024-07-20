package main

import "fmt"

func main() {
	fmt.Println("Top Down")
	fmt.Println(maximumScoreTopDown([]int{1, 2, 3}, []int{3, 2, 1}) == 14)
	fmt.Println(maximumScoreTopDown([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}) == 102)

	fmt.Println("Bottom Up")
	fmt.Println(maximumScoreTopDown([]int{1, 2, 3}, []int{3, 2, 1}) == maximumScoreBottomUp([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(maximumScoreTopDown([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}) == maximumScoreBottomUp([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}))
}

// Not sure how to memoize? Just look at the params we have l and m so see how we can memoize for those.

func maximumScoreTopDown(nums []int, multipliers []int) int {
	memo := make([][]int, len(multipliers))
	var dfs func(l, m int) int
	dfs = func(l, m int) int {
		if m >= len(multipliers) {
			return 0
		}

		if len(memo[m]) == 0 {
			memo[m] = make([]int, len(nums))
		}

		if memo[m][l] == 0 {
			r := len(nums) - 1 - (m - l)
			memo[m][l] = max(nums[l]*multipliers[m]+dfs(l+1, m+1), nums[r]*multipliers[m]+dfs(l, m+1))
		}

		return memo[m][l]
	}

	return dfs(0, 0)
}

func maximumScoreBottomUp(nums []int, multipliers []int) int {
	dp := make([][]int, len(multipliers)+1)
	for i := 0; i <= len(multipliers); i++ {
		dp[i] = make([]int, len(nums)+1)
	}

	for m := len(multipliers) - 1; m >= 0; m-- {
		for l := m; l >= 0; l-- {
			r := len(nums) - 1 - (m - l)
			dp[m][l] = max(nums[l]*multipliers[m]+dp[m+1][l+1], nums[r]*multipliers[m]+dp[m+1][l])
		}
	}

	return dp[0][0]
}
