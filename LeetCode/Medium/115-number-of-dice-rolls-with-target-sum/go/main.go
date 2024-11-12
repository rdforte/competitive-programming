package main

import "fmt"

func main() {
	fmt.Println(numRollsToTargetTopDown(1, 6, 3) == 1)
	fmt.Println(numRollsToTargetTopDown(2, 6, 7) == 6)
	fmt.Println(numRollsToTargetTopDown(30, 30, 500) == 222_616_187)
}

func numRollsToTarget(n, k, target int) int {
	var mod int = 1e9 + 7
	var dp func(diceNum, curSum int) int
	dp = func(diceNum, curSum int) int {
		if diceNum == n && curSum == 0 {
			return 1
		}

		if diceNum == n && curSum != 0 {
			return 0
		}

		total := 0
		for i := range k {
			total += dp(diceNum+1, curSum-(i+1))
		}

		return total % mod
	}

	return dp(0, target)
}

func numRollsToTargetTopDown(n, k, target int) int {
	var mod int = 1e9 + 7

	memo := make([][]int, 0, n)
	for range n {
		s := make([]int, 0, target+1)
		for range target + 1 {
			s = append(s, -1)
		}
		memo = append(memo, s)
	}

	var dp func(diceNum, curSum int) int
	dp = func(diceNum, curSum int) int {
		if diceNum == n {
			if curSum == 0 {
				return 1
			}
			return 0
		}

		if curSum < 0 {
			return 0
		}

		if memo[diceNum][curSum] == -1 {
			total := 0
			for i := range k {
				total += dp(diceNum+1, curSum-(i+1))
			}

			memo[diceNum][curSum] = total % mod
		}

		return memo[diceNum][curSum]
	}

	return dp(0, target)
}
