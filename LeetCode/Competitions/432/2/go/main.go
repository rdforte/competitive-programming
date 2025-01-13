package main

import (
	"fmt"
	"math"
)

func main() {
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
