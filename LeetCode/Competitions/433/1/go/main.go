package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{2, 3, 1}))    // 11
	fmt.Println(subarraySum([]int{3, 1, 1, 2})) // 13
	fmt.Println(subarraySum([]int{2, 5, 1, 5})) // 28
}

func subarraySum(nums []int) int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i] + prefix[i-1]
	}

	sum := 0

	for end := 0; end < len(nums); end++ {
		start := max(0, end-nums[end])

		prevSum := 0
		if start > 0 {
			prevSum = prefix[start-1]
		}
		sum += prefix[end] - prevSum
	}

	return sum
}
