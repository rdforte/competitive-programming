package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}) == 0)
	fmt.Println(maxProfit([]int{7, 1, 5, 0, 6, 4}) == 6)
	fmt.Println(maxProfit([]int{0, 1, 7, 0, 6, 4}) == 7)
}

func maxProfit(prices []int) int {
	var maxProfit float64

	for l, r := 0, 0; r < len(prices); r++ {
		maxProfit = math.Max(maxProfit, float64(prices[r]-prices[l]))
		if prices[r] < prices[l] {
			l = r
		}
	}

	return int(maxProfit)
}
