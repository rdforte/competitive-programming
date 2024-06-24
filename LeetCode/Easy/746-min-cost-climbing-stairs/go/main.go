package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCostClimbStairsTopDown([]int{10, 15, 20}) == minCostClimbStairsBottomUp([]int{10, 15, 20}))
	fmt.Println(minCostClimbStairsTopDown([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}) == minCostClimbStairsBottomUp([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbStairsTopDown(cost []int) int {
	cost = append(cost, 0) // Add 0 to represent the top
	memo := make(map[int]float64)

	var climb func(i int) float64
	climb = func(i int) float64 {
		if i < 0 {
			return 0
		}

		if _, ok := memo[i]; !ok {
			memo[i] = math.Min(float64(cost[i])+climb(i-1), float64(cost[i])+climb(i-2))
		}

		return memo[i]
	}

	return int(climb(len(cost) - 1))
}

func minCostClimbStairsBottomUp(cost []int) int {
	cost = append(cost, 0) // Add 0 to represent the top
	prev := cost[0]
	cur := cost[1]

	for i := 2; i < len(cost); i++ {
		tempCur := cur
		cur = int(math.Min(float64(cost[i]+cur), float64(cost[i]+prev)))
		prev = tempCur
	}

	return cur
}
