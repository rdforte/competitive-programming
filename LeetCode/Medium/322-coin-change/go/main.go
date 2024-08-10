package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("TOP DOWN")
	fmt.Println(coinChangeTopDown([]int{1, 2, 5}, 3) == 2)
	fmt.Println(coinChangeTopDown([]int{1, 2, 5}, 11) == 3)
	fmt.Println(coinChangeTopDown([]int{2}, 3) == -1)
	fmt.Println(coinChangeTopDown([]int{2}, 0) == 0)
	fmt.Println("BOTTOM UP")
	fmt.Println(coinChangeBottomUp([]int{1, 2, 5}, 3) == 2)
	fmt.Println(coinChangeBottomUp([]int{1, 2, 5}, 11) == 3)
	fmt.Println(coinChangeBottomUp([]int{2}, 3) == -1)
	fmt.Println(coinChangeBottomUp([]int{2}, 0) == 0)
}

func coinChangeTopDown(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -1
	}

	var dp func(amount int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}

		if amount < 0 {
			return math.MaxInt
		}

		if memo[amount] == -1 {
			minCoins := math.MaxInt
			for _, coin := range coins {
				if coinsUsed := dp(amount - coin); coinsUsed != math.MaxInt {
					minCoins = min(minCoins, coinsUsed+1)
				}
			}
			memo[amount] = minCoins
		}

		return memo[amount]
	}

	coinsUsed := dp(amount)
	if coinsUsed == math.MaxInt {
		return -1
	}

	return coinsUsed
}

func coinChangeBottomUp(coins []int, amount int) int {
	dp := make([]int, 0, amount+1)
	for range amount + 1 {
		dp = append(dp, math.MaxInt)
	}
	dp[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c < 0 || dp[a-c] == math.MaxInt {
				continue
			}
			dp[a] = min(dp[a], dp[a-c]+1)
		}
	}

	if dp[len(dp)-1] == math.MaxInt {
		return -1
	}

	return dp[len(dp)-1]
}
