package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}

// BRUTE FORCE
// Time Complexity - O(n^2)
func maxOperations(nums []int, k int) int {
	res := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == math.MaxInt32 {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j] == math.MaxInt32 {
				continue
			}

			if nums[i]+nums[j] == k {
				res++
				nums[i] = math.MaxInt32
				nums[j] = math.MaxInt32
				break
			}
		}
	}

	return res
}
