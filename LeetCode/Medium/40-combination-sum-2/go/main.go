package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var targetDFS func(nums []int, i, sum int)
	targetDFS = func(nums []int, i, sum int) {
		if sum > target {
			return
		}

		if sum == target {
			n := make([]int, len(nums))
			copy(n, nums)
			res = append(res, n)
			return
		}

		prev := -1
		for n := i; n < len(candidates); n++ {
			if candidates[n] == prev {
				continue
			}
			nums = append(nums, candidates[n])
			targetDFS(nums, n+1, sum+candidates[n])
			nums = nums[:len(nums)-1]
			prev = candidates[n]
		}
	}

	targetDFS([]int{}, 0, 0)

	return res
}
