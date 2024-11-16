package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		need := target - n
		if v, ok := seen[need]; ok {
			return []int{v, i}
		}
		seen[n] = i
	}

	return []int{}
}
