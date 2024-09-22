package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	maxP := 0         // havn't sold yet so we have 0 profit
	buy := -prices[0] // when we buy we are in the negative until we sell

	for i := 1; i < len(prices); i++ {
		buy = max(buy, -prices[i]) // we want to buy at the lowest price
		maxP = max(maxP, buy+prices[i])
	}

	return maxP
}
