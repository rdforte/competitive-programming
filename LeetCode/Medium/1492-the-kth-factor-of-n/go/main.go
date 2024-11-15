package main

import "fmt"

func main() {
	fmt.Println(kthFactor(12, 3))
	fmt.Println(kthFactor(7, 2))
	fmt.Println(kthFactor(4, 4))
	fmt.Println("--------------")
	fmt.Println(kthFactor2(12, 3))
	fmt.Println(kthFactor2(7, 2))
	fmt.Println(kthFactor2(4, 4))
}

// Brute Force, Loop through all of n
// Runtime is O(n), spacke O(1)
func kthFactor(n, k int) int {
	num := -1
	for i := 1; i <= n && k > 0; i++ {
		if n%i == 0 {
			num = i
			k--
		}
	}

	if k > 0 {
		return -1
	}

	return num
}

func kthFactor2(n, k int) int {
	num := -1
	for i := 1; i <= n/2 && k > 0; i++ {
		if n%i == 0 {
			num = i
			k--
		}
	}

	if k == 1 {
		return n
	}

	if k > 1 {
		return -1
	}

	return num
}
