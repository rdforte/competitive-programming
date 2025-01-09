package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
	fmt.Println(numOfSubarraysSimple([]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4))
}

// Overcomplicated by keeping track of avg
func numOfSubarrays(arr []int, k int, threshold int) int {
	res := 0
	avg := 0.0

	for l, r := 0, 0; r < len(arr); r++ {
		var rng float64 = float64((r - l) + 1)
		avg = ((avg * float64(rng-1)) + float64(arr[r])) / rng

		for rng > float64(k) {
			avg = ((avg * rng) - float64(arr[l])) / (rng - 1)
			rng--
			l++
		}

		if rng == float64(k) && avg >= float64(threshold) {
			res++
		}
	}

	return res
}

// Simplified solution keeping track of Sum
func numOfSubarraysSimple(arr []int, k int, threshold int) int {
	res, sum := 0, 0

	for l, r := 0, 0; r < len(arr); r++ {
		sum += arr[r]

		for (r-l)+1 > k {
			sum -= arr[l]
			l++
		}

		if (r-l)+1 == k && sum/k >= threshold {
			res++
		}
	}

	return res
}
