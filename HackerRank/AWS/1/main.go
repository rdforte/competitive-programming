package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxDiscountPairs([]int{2, 1, 8}))
	fmt.Println(getMaxDiscountPairs([]int{4, 5, 4, 5, 3, 6}))
}

// getting maxSum is O(n)
// generating powers of 3 is log3(maxSum) which can be O(log n)
// for each O(n) operation we do O(log n)
// there time complexity is O(n log n)

func getMaxDiscountPairs(price []int) int {
	var maxSum int
	for _, num := range price {
		maxSum = max(maxSum, num*2)
	}

	powers := generatePowersOfThree(maxSum)

	seen := make(map[int]int)
	seen[price[0]] = 1
	total := 0

	for i := 1; i < len(price); i++ {
		for p := range powers {
			diff := p - price[i]
			if v, ok := seen[diff]; ok {
				total += v
			}
		}
		seen[price[i]]++
	}

	return total
}

func generatePowersOfThree(limit int) map[int]struct{} {
	powers := make(map[int]struct{})
	powers[1] = struct{}{}
	for i := 3; i <= limit; i = i * 3 {
		powers[i] = struct{}{}
	}
	return powers
}
