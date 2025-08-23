package main

import (
	"fmt"
)

// https://neetcode.io/solutions/single-number-iii

func main() {
	res := singleNumbers([]int{1, 2, 1, 3, 2, 5})
	fmt.Println(res)
}

// Space O(1)
// Time O(n)
func singleNumbers(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	mask := 1
	for xor&mask != mask {
		mask <<= 1
	}

	group1, group2 := 0, 0

	for _, n := range nums {
		if n&mask == mask {
			group1 ^= n
		} else {
			group2 ^= n
		}
	}

	return []int{group1, group2}
}
