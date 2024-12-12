package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

// SPACE = O(n)
// TIME = O(n)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct, rightProduct := make([]int, n), make([]int, n)
	leftProduct[0], rightProduct[n-1] = nums[0], nums[n-1]

	for l, r := 1, n-2; l < n; l, r = l+1, r-1 {
		leftProduct[l] = leftProduct[l-1] * nums[l]
		rightProduct[r] = rightProduct[r+1] * nums[r]
	}

	if len(nums) > 1 {
		nums[0] = rightProduct[1]
		nums[len(nums)-1] = leftProduct[len(nums)-2]
	}

	for i := 1; i < len(nums)-1; i++ {
		nums[i] = leftProduct[i-1] * rightProduct[i+1]
	}

	return nums
}

// IMPROVED VERSION
// SPACE = O(1), disregard res as the output is not included in the space.
// TIME = O(n)
func productExceptSelfV2(nums []int) []int {
	n := len(nums)

	res := make([]int, n)
	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
