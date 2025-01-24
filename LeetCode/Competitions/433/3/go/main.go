package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCost(4, [][]int{
		{3, 5, 7},
		{6, 2, 9},
		{4, 8, 1},
		{7, 3, 5},
	}))
	fmt.Println(minCost(4, [][]int{
		{8, 2, 4},
		{6, 5, 4},
		{1, 9, 4},
		{6, 8, 7},
	}))
	fmt.Println(minCost(6, [][]int{
		{2, 4, 6},
		{5, 3, 8},
		{7, 1, 9},
		{4, 6, 2},
		{3, 5, 7},
		{8, 2, 4},
	}))
}

func minCost(n int, cost [][]int) int64 {
	numCosts := len(cost[0])

	memo := make([][]int, 0, n)
	for range n {
		costs := make([]int, 0, numCosts+1)
		for range numCosts + 1 {
			costs = append(costs, math.MaxInt64)
		}
		memo = append(memo, costs)
	}

	var dp func(h1, c int) int
	dp = func(h1, c int) int {
		if h1 >= n/2 {
			return 0
		}

		h2 := (n - 1) - h1

		if memo[h1][c] == math.MaxInt64 {
			for j := 1; j <= numCosts; j++ {
				if j == c {
					continue
				}
				price := dp(h1+1, j) + cost[h1][j-1]
				for k := 1; k <= numCosts; k++ {
					if j == k {
						continue
					}
					memo[h1][c] = min(memo[h1][c], price+cost[h2][k-1])
				}
			}
		}

		return memo[h1][c]
	}

	return int64(dp(0, 0))
}
