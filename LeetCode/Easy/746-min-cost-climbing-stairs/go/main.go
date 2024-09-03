package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbStairsTopDown([]int{10, 15, 20}) == minCostClimbStairsBottomUp([]int{10, 15, 20}))
	fmt.Println(minCostClimbStairsTopDown([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}) == minCostClimbStairsBottomUp([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbStairsTopDown(cost []int) int {
	cost = append(cost, 0) // Add 0 to represent the top
	memo := make(map[int]int)

	var climb func(i int) int
	climb = func(i int) int {
		if i < 0 {
			return 0
		}

		if _, ok := memo[i]; !ok {
			memo[i] = min(cost[i]+climb(i-1), cost[i]+climb(i-2))
		}

		return memo[i]
	}

	return climb(len(cost) - 1)
}

func minCostClimbStairsBottomUp(cost []int) int {
	cost = append(cost, 0) // Add 0 to represent the top
	prev, cur := cost[0], cost[1]

	for i := 2; i < len(cost); i++ {
		prev, cur = cur, min(cost[i]+cur, cost[i]+prev)
	}

	return cur
}
