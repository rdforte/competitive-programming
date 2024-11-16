package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	seen := make([]int, 0, len(nums))
	res := make([][]int, 0)
	dupl := make(map[[3]int]struct{})

	for _, n := range nums {
		target := 0 - n

		prev := make(map[int]struct{})
		for _, s := range seen {
			need := target - s
			if _, ok := prev[need]; ok {
				ans := []int{n, s, need}
				slices.Sort(ans)
				key := [3]int{ans[0], ans[1], ans[2]}
				if _, isDupl := dupl[key]; !isDupl {
					res = append(res, ans)
					dupl[key] = struct{}{}
				}
			}
			prev[s] = struct{}{}
		}

		seen = append(seen, n)
	}

	return res
}
