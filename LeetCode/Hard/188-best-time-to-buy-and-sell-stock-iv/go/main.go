package main

import "fmt"

func main() {
	fmt.Println("TOP DOWN")
	fmt.Println(maxProfitTopDown(2, []int{2, 4, 1}))
	fmt.Println(maxProfitTopDown(2, []int{3, 2, 6, 5, 0, 3}))
	fmt.Println("BOTTOM UP")
	fmt.Println(maxProfitBottomUp(2, []int{2, 4, 1}))
	fmt.Println(maxProfitBottomUp(2, []int{3, 2, 6, 5, 0, 3}))
}

// At each price (p) we have a few decisions based on whether we are holding
// stock or not holding stock.
// If we are NOT holding stock then we can decide to `doNothing` or `buyStock`.
// If we are holding stock then we can decide to `doNothing` or `sellStock`.
// We can repeat this process recursively to find the max profit with the max number of transactions allowed (k).

// note that each function call takes three arguments p (prices), k (numTransactions), h (holding stock) because we need to keep track
// of this state throughout our recursion.
// So to memoize this function call we need a slice which is 3D so we can reference each index based on the params passed to the function.

func maxProfitTopDown(numTrans int, prices []int) int {
	var memo [][][]int
	for range len(prices) {
		m := make([][]int, 0, numTrans)
		for range numTrans {
			h := make([]int, 0, 2)
			h = append(h, -1, -1)
			m = append(m, h)
		}
		memo = append(memo, m)
	}

	var dp func(p, k, h int) int
	dp = func(p, k, h int) int {
		if p >= len(prices) || k < 0 {
			return 0
		}

		if memo[p][k][h] == -1 {
			profit := 0
			if h == 0 {
				profit = max(dp(p+1, k, h), -prices[p]+dp(p+1, k, 1))
			}

			if h == 1 {
				profit = max(dp(p+1, k, h), prices[p]+dp(p+1, k-1, 0))
			}

			memo[p][k][h] = profit
		}

		return memo[p][k][h]
	}

	return dp(0, numTrans-1, 0)
}

// With Bottomo up remember we start at the base case

func maxProfitBottomUp(numTrans int, prices []int) int {
	var dp [][][]int
	// +1 to account for base case
	for range len(prices) + 1 {
		d := make([][]int, 0, numTrans+1)
		for range numTrans + 1 {
			h := make([]int, 0, 2)
			h = append(h, 0, 0)
			d = append(d, h)
		}
		dp = append(dp, d)
	}

	for p := len(prices) - 1; p >= 0; p-- {
		for k := 1; k <= numTrans; k++ {
			for h := 0; h < 2; h++ {
				profit := 0
				if h == 0 {
					profit = max(dp[p+1][k][h], -prices[p]+dp[p+1][k][1])
				}

				if h == 1 {
					profit = max(dp[p+1][k][h], prices[p]+dp[p+1][k-1][0])
				}

				dp[p][k][h] = profit
			}
		}
	}

	return dp[0][numTrans][0]
}
