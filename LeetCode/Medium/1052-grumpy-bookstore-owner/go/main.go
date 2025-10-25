package main

import "math"

func main() {
	println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	println(maxSatisfied([]int{1}, []int{0}, 1))
	println(maxSatisfied([]int{4, 10, 10}, []int{1, 1, 0}, 2))
}

// we want to schedule this window during the time period that would save the
// largest number of customers from the bookstore owners grumpiness.

// space = O(1)
// time = O(2n) = O(n). We have to do 2 passes of the array
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	maxCustomerWithGrumpy, curCustomerWithGrumpy := math.MinInt, 0
	leftI, rightI := 0, 0 // keep the range for which me can max customers which would have been grumpy.

	for l, r := 0, 0; r < len(customers); r++ {
		if grumpy[r] == 1 {
			curCustomerWithGrumpy += customers[r]
		}
		if (r+1)-l > minutes {
			if grumpy[l] == 1 {
				curCustomerWithGrumpy -= customers[l]
			}
			l++
		}
		if (r+1)-l == minutes && curCustomerWithGrumpy > maxCustomerWithGrumpy {
			maxCustomerWithGrumpy = curCustomerWithGrumpy
			leftI, rightI = l, r
		}
	}

	res := 0

	for i, customer := range customers {
		inRange := i >= leftI && i <= rightI
		if inRange {
			res += customer
			continue
		}
		if grumpy[i] == 0 {
			res += customer
		}
	}

	return res
}
