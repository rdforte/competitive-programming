package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	nums = []int{-1}
	rotate(nums, 2)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k %= len(nums) // if k is larger than len(nums) then this will give us the remainder which we need to shift.

	newNums := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i, n := range newNums {
		nums[i] = n
	}
}
