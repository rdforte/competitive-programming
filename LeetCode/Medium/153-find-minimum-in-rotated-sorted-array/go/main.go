package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{3, 1, 2}))
}

// One thing to note here is that when we do nums[r] < mid as in is right less than mid then we know that mid cant be our answer so we can set left to mid +1.
// but that also means if we shift our right pointer to mid mid can be inclusive of the answer we are looking for so we don't do r = mid -1 and just r = mid

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		midI := l + (r-l)/2
		mid := nums[midI]

		if nums[r] < mid {
			l = midI + 1
		} else {
			r = midI
		}
	}
	return nums[l]
}
