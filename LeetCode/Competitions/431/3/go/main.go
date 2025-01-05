package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumCoins([][]int{
		{8, 10, 1},
		{1, 3, 2},
		{5, 6, 4},
	}, 4))
	fmt.Println(maximumCoins([][]int{
		{1, 10, 3},
	}, 2))
}

func maximumCoins(coins [][]int, k int) int64 {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i][0] < coins[j][0]
	})

	nums := make([]int, coins[len(coins)-1][1]+1)

	for _, c := range coins {
		for i := c[0]; i <= c[1]; i++ {
			nums[i] = c[2]
		}
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	if k >= len(nums) {
		return int64(nums[len(nums)-1])
	}

	res := nums[k]

	for i := k + 1; i < len(nums); i++ {
		res = max(res, nums[i]-nums[i-k])
	}

	return int64(res)
}
