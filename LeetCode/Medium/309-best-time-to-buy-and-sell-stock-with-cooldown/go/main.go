package main

import "fmt"

func main() {
	fmt.Println("TOP DOWN")
	fmt.Println(maxProfitTopDown([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfitTopDown([]int{1}))
	fmt.Println("BOTTOM UP")
	fmt.Println(maxProfitBottomUp([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfitBottomUp([]int{1}))
}

// Brute Force recursion
func maxProfit(prices []int) int {
	var dp func(p, c, h int) int
	dp = func(p, c, h int) int {
		if p >= len(prices) {
			return 0
		}

		// has sold but can't buy because of cooldown
		maxProfit := dp(p+1, 0, h)

		if h == 1 {
			// has baught so can sell
			maxProfit = max(maxProfit, prices[p]+dp(p+1, 1, 0))
		}

		// can buy if no cooldown
		if c == 0 && h == 0 {
			maxProfit = max(maxProfit, -prices[p]+dp(p+1, 0, 1))
		}

		return maxProfit
	}

	return dp(0, 0, 0)
}

// Recursion with memoisation
func maxProfitTopDown(prices []int) int {
	memo := make([][][]int, 0, len(prices))
	for range len(prices) {
		cooldown := make([][]int, 0, 2)
		for range 2 {
			holding := make([]int, 0, 2)
			holding = append(holding, -1, -1)
			cooldown = append(cooldown, holding)
		}
		memo = append(memo, cooldown)
	}

	var dp func(p, c, h int) int
	dp = func(p, c, h int) int {
		if p >= len(prices) {
			return 0
		}

		if memo[p][c][h] == -1 {
			// has sold but can't buy because of cooldown
			maxProfit := dp(p+1, 0, h)

			if h == 1 {
				// has baught so can sell
				maxProfit = max(maxProfit, prices[p]+dp(p+1, 1, 0))
			}

			// can buy if no cooldown
			if c == 0 && h == 0 {
				maxProfit = max(maxProfit, -prices[p]+dp(p+1, 0, 1))
			}

			memo[p][c][h] = maxProfit
		}

		return memo[p][c][h]
	}

	return dp(0, 0, 0)
}

// Bottom Up recursive. Start at base case
func maxProfitBottomUp(prices []int) int {
	dp := make([][][]int, 0, len(prices)+1) // add 1 for base case
	for range len(prices) + 1 {
		cooldown := make([][]int, 0, 2)
		for range 2 {
			holding := make([]int, 2)
			cooldown = append(cooldown, holding)
		}
		dp = append(dp, cooldown)
	}

	// base case is already set so start iterating at the 1 after the base case
	for p := len(prices) - 1; p >= 0; p-- {
		for c := 0; c < 2; c++ {
			for h := 0; h < 2; h++ {
				// has sold but can't buy because of cooldown
				maxProfit := dp[p+1][0][h]

				if h == 1 {
					// has baught so can sell
					maxProfit = max(maxProfit, prices[p]+dp[p+1][1][0])
				}

				// can buy if no cooldown
				if c == 0 && h == 0 {
					maxProfit = max(maxProfit, -prices[p]+dp[p+1][0][1])
				}

				dp[p][c][h] = maxProfit
			}
		}
	}

	return dp[0][0][0]
}
