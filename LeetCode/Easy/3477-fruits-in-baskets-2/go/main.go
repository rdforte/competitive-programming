package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfUnplacedFruits([]int{4, 2, 5}, []int{3, 5, 4}))
	fmt.Println(numOfUnplacedFruits([]int{3, 6, 1}, []int{6, 4, 7}))
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	res := 0

LOOP:
	for _, fruit := range fruits {
		for b, basket := range baskets {
			if fruit <= basket {
				baskets[b] = 0
				continue LOOP
			}
		}
		res++
	}

	return res
}
