package main

import "fmt"

func main() {
	fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 3}))
	fmt.Println(maximumPrimeDifference([]int{4, 8, 2, 8}))
	fmt.Println(maximumPrimeDifference([]int{4, 8, 2, 8}))
}

func maximumPrimeDifference(nums []int) int {
	furthestRightPrime := 0
	furthestLeftPrime := 0

	for l, r := 0, 0; r < len(nums); r++ {
		if !isPrime(nums[l]) {
			l++
			continue
		} else {
			furthestLeftPrime = l
		}

		if isPrime(nums[r]) {
			furthestRightPrime = r
		}
	}

	return furthestRightPrime - furthestLeftPrime
}

func isPrime(n int) bool {
	prime := map[int]struct{}{
		2:  {},
		3:  {},
		5:  {},
		7:  {},
		11: {},
		13: {},
		17: {},
		19: {},
		23: {},
		29: {},
		31: {},
		37: {},
		41: {},
		43: {},
		47: {},
		53: {},
		59: {},
		61: {},
		67: {},
		71: {},
		73: {},
		79: {},
		83: {},
		89: {},
		97: {},
	}

	if _, ok := prime[n]; ok {
		return true
	}

	return false
}
