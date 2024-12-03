package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	// -5,-5,-4,0,0,3,3,4,5
	fmt.Println(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
}

// time complexity =
// n.log.n to sort plus n^2 to find target sum.
// n.log.n + n^2 =
// O(n^2)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	numsLen := len(nums)
	res := math.MaxInt

	for i := 0; i < numsLen-2; i++ {
		for l, r := i+1, numsLen-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]

			if abs(target-sum) < abs(target-res) {
				res = sum
			}

			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
