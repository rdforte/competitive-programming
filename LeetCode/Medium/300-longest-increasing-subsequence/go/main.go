package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{7, 7}))
	fmt.Println(lengthOfLIS([]int{7, 7, 8}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}

	return res
}
