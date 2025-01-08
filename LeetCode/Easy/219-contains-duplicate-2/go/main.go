package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{0, 1, 2, 3, 2, 5}, 3))
}

// First Approach
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	windowItems := make(map[int]int)

	for l, r := 0, 0; r < len(nums); r++ {

		for r-l > k {
			windowItems[nums[l]]--
			l++
		}

		if v, ok := windowItems[nums[r]]; ok && v > 0 {
			return true
		}

		windowItems[nums[r]]++
	}

	return false
}

// Second Approach
// Keep track of pointer
func containsNearbyDuplicateImproved(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	windowItems := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := windowItems[nums[i]]; ok && i-v <= k {
			return true
		}
		windowItems[nums[i]] = i // override the prev occurence of this num as it will be the closest to the idx.
	}

	return false
}
