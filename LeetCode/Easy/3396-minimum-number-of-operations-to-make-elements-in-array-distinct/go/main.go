package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4, 2, 3, 3, 5, 7}))
	fmt.Println(minimumOperations([]int{4, 5, 6, 4, 4}))
	fmt.Println(minimumOperations([]int{6, 7, 8, 9}))
}

func minimumOperations(nums []int) int {
	mp := make(map[int]int, len(nums))

	unique, repeated, res := 0, len(nums), 0

	for _, n := range nums {
		if _, ok := mp[n]; !ok {
			unique++
		}
		mp[n]++
	}

	i := 0
	for unique != repeated {
		repeated--
		n := nums[i]
		count := mp[n]

		if count == 1 {
			unique--
		}

		mp[n]--
		if i%3 == 0 {
			res++
		}
		i++
	}

	return res
}
