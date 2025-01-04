package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))

	fmt.Println(maxOperationsOptimised2Pass([]int{1, 2, 3, 4}, 5))
	fmt.Println(maxOperationsOptimised2Pass([]int{3, 1, 3, 4, 3}, 6))
	fmt.Println(maxOperationsOptimised2Pass([]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2))
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

// Space O(n)
// time O(n)
// This is similar to 2Sum in the sense that we can see the numbers we have
// and then find the compliment.
func maxOperationsOptimised2Pass(nums []int, k int) int {
	mp := make(map[int]int, len(nums))

	for _, n := range nums {
		mp[n]++
	}

	res := 0
	for _, n := range nums {
		need := k - n

		cur, hasCur := mp[n]
		comp, hasComp := mp[need]

		if hasCur && hasComp && cur > 0 && comp > 0 {
			if n == need && cur < 2 {
				continue
			}
			mp[need]--
			mp[n]--
			res++
		}

	}

	return res
}
