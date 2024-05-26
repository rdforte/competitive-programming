package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3))
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{}, 3))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	res := 0

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				res++
			}
		}
	}

	return res
}
