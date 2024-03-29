package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		seen := make(map[int]struct{})
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			sKey := 0 - nums[i] - nums[j]
			if _, ok := seen[sKey]; ok {
				res = append(res, []int{nums[i], sKey, nums[j]})
				nextJ := j
				for nextJ+1 < len(nums) && nums[nextJ] == nums[nextJ+1] {
					nextJ++
					j = nextJ
				}
			}
			seen[nums[j]] = struct{}{}
		}
	}

	return res
}
