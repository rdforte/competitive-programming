package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("V1 ----------------------------------------")
	fmt.Println(maximumAmount([][]int{
		{0, 1, -1},
		{1, -2, 3},
		{2, -3, 4},
	}))
	fmt.Println(maximumAmount([][]int{
		{-7, 12, 12, 13},
		{-6, 19, 19, -6},
		{9, -2, -10, 16},
		{-4, 14, -10, -9},
	})) // 60

	fmt.Println(maximumAmount([][]int{
		{-16, 4, 1, -1},
		{11, 9, 3, 3},
		{-6, 17, -19, 9},
		{14, -17, -19, -13},
	})) // 35
	fmt.Println("V2 ----------------------------------------")
	fmt.Println(maximumAmountV2([][]int{
		{0, 1, -1},
		{1, -2, 3},
		{2, -3, 4},
	}))
	fmt.Println(maximumAmountV2([][]int{
		{-7, 12, 12, 13},
		{-6, 19, 19, -6},
		{9, -2, -10, 16},
		{-4, 14, -10, -9},
	})) // 60

	fmt.Println(maximumAmountV2([][]int{
		{-16, 4, 1, -1},
		{11, 9, 3, 3},
		{-6, 17, -19, 9},
		{14, -17, -19, -13},
	})) // 35
	fmt.Println("V3 ----------------------------------------")
	fmt.Println(maximumAmountV3BottomUp([][]int{
		{0, 1, -1},
		{1, -2, 3},
		{2, -3, 4},
	}))
	fmt.Println(maximumAmountV3BottomUp([][]int{
		{-7, 12, 12, 13},
		{-6, 19, 19, -6},
		{9, -2, -10, 16},
		{-4, 14, -10, -9},
	})) // 60

	fmt.Println(maximumAmountV3BottomUp([][]int{
		{-16, 4, 1, -1},
		{11, 9, 3, 3},
		{-6, 17, -19, 9},
		{14, -17, -19, -13},
	})) // 35

	fmt.Println(maximumAmountV3BottomUp([][]int{
		{-6, -15, -16, -8},
		{-10, 11, 6, 16},
		{1, 2, 18, 12},
		{15, 19, 4, 17},
	})) // 64
}

func maximumAmount(coins [][]int) int {
	memo := make([][][]int, 0, len(coins))
	for range len(coins) {
		c := make([][]int, 0, len(coins[0]))
		for range len(coins[0]) {
			neutralised := make([]int, 3)
			neutralised[0], neutralised[1], neutralised[2] = math.MinInt64, math.MinInt64, math.MinInt64
			c = append(c, neutralised)
		}
		memo = append(memo, c)
	}

	var dp func(r, c, neutralised int) int
	dp = func(r, c, neutralised int) int {
		if r == len(coins)-1 && c == len(coins[r])-1 {
			if coins[r][c] < 0 && neutralised < 2 {
				return 0
			}
			return coins[r][c]
		}

		if memo[r][c][neutralised] == math.MinInt64 {
			sum := math.MinInt64
			if r+1 < len(coins) {
				sum = dp(r+1, c, neutralised) + coins[r][c]
			}
			if c+1 < len(coins[r]) {
				sum = max(sum, dp(r, c+1, neutralised)+coins[r][c])
			}
			if coins[r][c] < 0 && neutralised < 2 {
				if r+1 < len(coins) {
					sum = max(sum, dp(r+1, c, neutralised+1))
				}
				if c+1 < len(coins[r]) {
					sum = max(sum, dp(r, c+1, neutralised+1))
				}
			}
			memo[r][c][neutralised] = sum
		}

		return memo[r][c][neutralised]
	}

	return dp(0, 0, 0)
}

func maximumAmountV2(coins [][]int) int {
	memo := make([][][]int, 0, len(coins))
	for range len(coins) {
		c := make([][]int, 0, len(coins[0]))
		for range len(coins[0]) {
			neutralised := make([]int, 3)
			neutralised[0], neutralised[1], neutralised[2] = math.MinInt64, math.MinInt64, math.MinInt64
			c = append(c, neutralised)
		}
		memo = append(memo, c)
	}

	var dp func(r, c, neutralised int) int
	dp = func(r, c, neutralised int) int {
		if r >= len(coins) || c >= len(coins[r]) {
			return math.MinInt32 // return -Int32 as -Int64 can cause overflow
		}

		if r == len(coins)-1 && c == len(coins[r])-1 {
			if coins[r][c] < 0 && neutralised < 2 {
				return 0
			}
			return coins[r][c]
		}

		if memo[r][c][neutralised] == math.MinInt64 {
			sum := max(dp(r+1, c, neutralised)+coins[r][c], dp(r, c+1, neutralised)+coins[r][c])
			if coins[r][c] < 0 && neutralised < 2 {
				sum = max(sum, dp(r+1, c, neutralised+1), dp(r, c+1, neutralised+1))
			}
			memo[r][c][neutralised] = sum
		}

		return memo[r][c][neutralised]
	}

	return dp(0, 0, 0)
}

// With bottom up we always start with the bottom case. This will be the base case from iterative solution.
// In the iterative solution we can step outside the bounds of the grid in which case we return MIN_INT.
// We will initialise our 3D DP array to have +1 rows and cols with MIN_INT.
func maximumAmountV3BottomUp(coins [][]int) int {
	dp := make([][][]int, 0, len(coins)+1)
	for range len(coins) + 1 {
		c := make([][]int, 0, len(coins[0])+1)
		for range len(coins[0]) + 1 {
			neutralised := make([]int, 3)
			neutralised[0], neutralised[1], neutralised[2] = math.MinInt32, math.MinInt32, math.MinInt32
			c = append(c, neutralised)
		}
		dp = append(dp, c)
	}

	for r := len(coins) - 1; r >= 0; r-- {
		for c := len(coins[r]) - 1; c >= 0; c-- {
			for neutralised := 2; neutralised >= 0; neutralised-- {
				if r == len(coins)-1 && c == len(coins[r])-1 {
					if coins[r][c] < 0 && neutralised < 2 {
						dp[r][c][neutralised] = 0
					} else {
						dp[r][c][neutralised] = coins[r][c]
					}
					continue
				}
				sum := max(dp[r+1][c][neutralised]+coins[r][c], dp[r][c+1][neutralised]+coins[r][c])
				if coins[r][c] < 0 && neutralised < 2 {
					sum = max(sum, dp[r+1][c][neutralised+1], dp[r][c+1][neutralised+1])
				}
				dp[r][c][neutralised] = sum
			}
		}
	}

	return max(dp[0][0][0], dp[0][0][1], dp[0][0][2])
}
