package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{3, 3, 3, 3, 3}))
}

// Linked List
// Floyd's algorithm - used for finding the start of a cycle

func findDuplicate(nums []int) int {
	// Because we want to keep looping until both pointers are the same and if we start at 0 for both they will be the same,
	// so lets start by moving slow forward one and fast forward 2 before looping
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
	}

	slow2 := 0
	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow
}
