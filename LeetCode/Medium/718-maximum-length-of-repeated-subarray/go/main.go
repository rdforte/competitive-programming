package main

import "fmt"

func main() {
	fmt.Println(findLength([]int{1, 2}, []int{1, 2}) == 2)
	fmt.Println(findLength([]int{0, 1, 2}, []int{3, 1, 2}) == 2)
	fmt.Println(findLength([]int{0, 1, 1}, []int{3, 1, 4}) == 1)
	fmt.Println(findLength([]int{0, 1, 2}, []int{3, 1, 2, 2}) == 2)
	fmt.Println(findLength([]int{0, 1, 2, 2}, []int{3, 1, 2}) == 2)
	fmt.Println(findLength([]int{1, 2}, []int{3, 4}) == 0)
	fmt.Println(findLength([]int{1, 2}, []int{2, 1}) == 1)
	fmt.Println(findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}) == 2)
}

func findLength(nums1, nums2 []int) int {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	memo := make([][]int, 0, nums1Len+1)
	for range nums1Len + 1 {
		memo = append(memo, make([]int, nums2Len+1))
	}

	res := 0
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				memo[i][j] = memo[i+1][j+1] + 1
			}
			res = max(res, memo[i][j])
		}
	}

	return res
}
