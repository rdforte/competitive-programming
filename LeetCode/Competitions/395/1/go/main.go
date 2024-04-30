package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(addedInteger([]int{2, 6, 4}, []int{9, 7, 5}))
	fmt.Println(addedInteger([]int{10}, []int{5}))
	fmt.Println(addedInteger([]int{1, 1, 1}, []int{1, 1, 1}))
}

func addedInteger(nums1 []int, nums2 []int) int {
	s1 := nums1[0]
	s2 := nums2[0]

	for i := 1; i < len(nums1); i++ {
		s1 = int(math.Min(float64(s1), float64(nums1[i])))
		s2 = int(math.Min(float64(s2), float64(nums2[i])))
	}

	return int(s2 - s1)
}
