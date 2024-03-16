package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{3, 1, 2}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		midI := l + (r-l)/2
		mid := nums[midI]

		if nums[r] > mid {
			r = midI
		} else {
			l = midI + 1
		}
	}
	return nums[l]
}
