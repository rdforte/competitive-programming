package main

import "fmt"

func main() {
	fmt.Println(countAlternatingSubarrays([]int{1, 0, 1, 0}))
	fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1}))
}

func countAlternatingSubarrays(nums []int) int64 {
	res, size := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			size = 1
			res = res + size
			continue
		}
		size++
		res = res + size
	}

	return int64(res)
}
