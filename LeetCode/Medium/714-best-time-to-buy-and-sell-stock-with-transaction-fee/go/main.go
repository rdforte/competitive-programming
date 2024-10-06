package main

import "fmt"

func main() {
	fmt.Println("TOP DOWN")
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2) == maxProfitTopDown([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3) == maxProfitTopDown([]int{1, 3, 7, 5, 10, 3}, 3))
	fmt.Println("BOTTOM UP")
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2) == maxProfitBottomUp([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3) == maxProfitBottomUp([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	var dp func(i int, holding bool) int
	dp = func(i int, holding bool) int {
		if i >= len(prices) {
			return 0
		}

		profit := dp(i+1, holding)

		if !holding {
			profit = max(profit, -prices[i]+dp(i+1, true))
		}

		if holding {
			profit = max(profit, (prices[i]-fee)+dp(i+1, false))
		}

		return profit
	}

	return dp(0, false)
}

func maxProfitTopDown(prices []int, fee int) int {
	nPrices := len(prices)
	memo := make([][]int, 0, nPrices)
	for range nPrices {
		memoHolding := []int{-1, -1}
		memo = append(memo, memoHolding)
	}

	var dp func(i int, holding int) int
	dp = func(i int, holding int) int {
		if i >= len(prices) {
			return 0
		}

		if memo[i][holding] == -1 {

			profit := dp(i+1, holding)

			if holding == 0 {
				profit = max(profit, -prices[i]+dp(i+1, 1))
			}

			if holding == 1 {
				profit = max(profit, (prices[i]-fee)+dp(i+1, 0))
			}
			memo[i][holding] = profit
		}

		return memo[i][holding]
	}

	return dp(0, 0)
}

// NOTES:
// if we look at maxProfitTopDown we can see the base case is when p >= len(prices) in which case we return 0.
// we can take this information to then start at our base case.
// our max profit will also be when we are not holding any stocks and we have bought and sold in which case holding = 0.
func maxProfitBottomUp(prices []int, fee int) int {
	nPrices := len(prices)
	dp := make([][]int, 0, nPrices+1)
	for range nPrices + 1 {
		memoHolding := []int{0, 0}
		dp = append(dp, memoHolding)
	}

	for p := nPrices - 1; p >= 0; p-- {
		for h := 1; h >= 0; h-- {
			profit := dp[p+1][h]

			if h == 0 {
				profit = max(profit, -prices[p]+dp[p+1][1])
			}

			if h == 1 {
				profit = max(profit, (prices[p]-fee)+dp[p+1][0])
			}
			dp[p][h] = profit
		}
	}

	return dp[0][0]
}
