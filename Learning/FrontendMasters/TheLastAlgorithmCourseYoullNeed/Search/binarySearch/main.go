package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 5))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 6, 7, 8}, 5))

	fmt.Println(twoCrystalBalls([]bool{false, false, false, false, false, false, false, false, false, true, true}))
}

func binarySearch(s []int, k int) int {
	low := 0
	high := len(s) - 1

	for low <= high {
		mid := low + (high-low)/2 // (low + high) / 2 is a famous bug and can cause overflow
		val := s[mid]

		// the mid point is not k so then it must be on either side
		if val == k {
			return mid
		} else if val > k {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1 // sentinal value
}

// time complexity O(√n)
// we jump forward √n times then move back √n and walk forward n times till we find the first true
func twoCrystalBalls(breaks []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(breaks))))

	i := 0
	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount // go back to where we didnt have break so can walk forward to nearest breaks

	for j := i; j < len(breaks); j++ {
		if breaks[j] {
			return j
		}
	}

	return -1
}
