package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4, 4}))
}

func deleteAndEarn(nums []int) int {
	mx := slices.Max(nums)

	group := make([]int, mx+1)
	for _, n := range nums {
		group[n] += n
	}

	if len(group) >= 2 {
		group[1] = max(group[0], group[1])
	}

	for i := 2; i < len(group); i++ {
		group[i] = max(group[i-1], group[i]+group[i-2])
	}

	return group[len(group)-1]
}
