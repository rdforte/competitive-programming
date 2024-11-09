package main

import "fmt"

func main() {
	fmt.Println("BRUTE FORCE RECURSION")
	fmt.Println(findLength([]int{1, 2}, []int{1, 2}) == 2)
	fmt.Println(findLength([]int{0, 1, 2}, []int{3, 1, 2}) == 2)
	fmt.Println(findLength([]int{0, 1, 1}, []int{3, 1, 4}) == 1)
	fmt.Println(findLength([]int{0, 1, 2}, []int{3, 1, 2, 2}) == 2)
	fmt.Println(findLength([]int{0, 1, 2, 2}, []int{3, 1, 2}) == 2)
	fmt.Println(findLength([]int{1, 2}, []int{3, 4}) == 0)
	fmt.Println(findLength([]int{1, 2}, []int{2, 1}) == 1)
	fmt.Println("DP - TOP DOWN")
	fmt.Println(findLengthTopDown([]int{1, 2}, []int{1, 2}) == 2)
	fmt.Println(findLengthTopDown([]int{0, 1, 2}, []int{3, 1, 2}) == 2)
	fmt.Println(findLengthTopDown([]int{0, 1, 1}, []int{3, 1, 4}) == 1)
	fmt.Println(findLengthTopDown([]int{0, 1, 2}, []int{3, 1, 2, 2}) == 2)
	fmt.Println(findLengthTopDown([]int{0, 1, 2, 2}, []int{3, 1, 2}) == 2)
	fmt.Println(findLengthTopDown([]int{1, 2}, []int{3, 4}) == 0)
	fmt.Println(findLengthTopDown([]int{1, 2}, []int{2, 1}) == 1)

	fmt.Println(findLengthTopDown([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1})) // 2
}

func findLength(nums1, nums2 []int) int {
	var dp func(n1, n2 int) int
	dp = func(n1, n2 int) int {
		if n1 >= len(nums1) || n2 >= len(nums2) {
			return 0
		}

		maxLen := 0

		if nums1[n1] == nums2[n2] {
			maxLen = max(maxLen, dp(n1+1, n2+1)+1)
		}

		maxLen = max(maxLen, dp(n1, n2+1))
		maxLen = max(maxLen, dp(n1+1, n2))

		return maxLen
	}

	return dp(0, 0)
}

func findLengthTopDown(nums1, nums2 []int) int {
	memo := make([][]int, 0, len(nums1))
	for range len(nums1) {
		n2 := make([]int, 0, len(nums2))
		for range len(nums2) {
			n2 = append(n2, -1)
		}
		memo = append(memo, n2)
	}

	var dp func(n1, n2 int) int
	dp = func(n1, n2 int) int {
		if n1 >= len(nums1) || n2 >= len(nums2) {
			return 0
		}

		if memo[n1][n2] == -1 {
			maxLen := 0

			if nums1[n1] == nums2[n2] {
				maxLen = max(maxLen, dp(n1+1, n2+1)+1)
			}

			maxLen = max(maxLen, dp(n1, n2+1))
			maxLen = max(maxLen, dp(n1+1, n2))

			memo[n1][n2] = maxLen
		}

		return memo[n1][n2]
	}

	return dp(0, 0)
}
