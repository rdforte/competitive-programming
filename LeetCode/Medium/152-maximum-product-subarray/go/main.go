package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	res := float64(nums[0])
	dp := make([][]float64, len(nums)) // [max, min]
	dp[0] = []float64{float64(nums[0]), float64(nums[0])}

	for i := 1; i < len(nums); i++ {
		cur := float64(nums[i])

		if cur == 0 {
			dp[i] = []float64{0, 0}
			res = math.Max(res, 0)
			continue
		}

		prevMax := dp[i-1][0]
		prevMin := dp[i-1][1]

		max := math.Max(math.Max(cur, cur*prevMax), cur*prevMin)
		min := math.Min(math.Min(cur, cur*prevMax), cur*prevMin)

		dp[i] = []float64{max, min}

		res = math.Max(res, max)
	}

	return int(res)
}
