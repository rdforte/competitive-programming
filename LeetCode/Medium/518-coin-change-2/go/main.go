package main

import "fmt"

func main() {
	fmt.Println("BRUTE FORCE")
	fmt.Println(changeBruteForce(5, []int{1, 2, 5}))
	fmt.Println(changeBruteForce(3, []int{2}))
	fmt.Println(changeBruteForce(10, []int{1}))
	fmt.Println("TOP DOWN")
	fmt.Println(changeTopDown(5, []int{1, 2, 5}))
	fmt.Println(changeTopDown(3, []int{2}))
	fmt.Println(changeTopDown(10, []int{1}))
	fmt.Println("BOTTOM UP")
	fmt.Println(changeBottomUp(5, []int{1, 2, 5}))
	fmt.Println(changeBottomUp(3, []int{2}))
	fmt.Println(changeBottomUp(10, []int{1}))
}

func changeBruteForce(amount int, coins []int) int {
	var dp func(i, remainingAmount int) int
	dp = func(i, remainingAmount int) int {
		if i >= len(coins) {
			return 0
		}

		if remainingAmount == 0 {
			return 1
		}

		if remainingAmount < 0 {
			return 0
		}

		combinations := dp(i, remainingAmount-coins[i]) + dp(i+1, remainingAmount)

		return combinations
	}

	return dp(0, amount)
}

func changeTopDown(amount int, coins []int) int {
	memo := make([][]int, 0, len(coins))
	for range len(coins) {
		m := make([]int, 0, amount+1)
		for range amount + 1 {
			m = append(m, -1)
		}
		memo = append(memo, m)
	}

	var dp func(i, remainingAmount int) int
	dp = func(i, remainingAmount int) int {
		if i >= len(coins) {
			return 0
		}

		if remainingAmount == 0 {
			return 1
		}

		if remainingAmount < 0 {
			return 0
		}

		if memo[i][remainingAmount] == -1 {
			memo[i][remainingAmount] = dp(i, remainingAmount-coins[i]) + dp(i+1, remainingAmount)
		}

		return memo[i][remainingAmount]
	}

	return dp(0, amount)
}

func changeBottomUp(amount int, coins []int) int {
	dp := make([][]int, 0, len(coins)+1)
	for range len(coins) + 1 {
		dp = append(dp, make([]int, amount+1))
	}
	for i := range dp {
		dp[i][0] = 1
	}

	for c := len(coins) - 1; c >= 0; c-- {
		for a := 1; a <= amount; a++ {
			if a-coins[c] < 0 {
				dp[c][a] = dp[c+1][a]
			} else {
				dp[c][a] = dp[c][a-coins[c]] + dp[c+1][a]
			}
		}
	}

	return dp[0][amount]
}
