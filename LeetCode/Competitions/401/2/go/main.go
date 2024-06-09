package main

import "fmt"

func main() {
	fmt.Println(valueAfterKSeconds(4, 5))
	fmt.Println(valueAfterKSeconds(5, 3))
	fmt.Println(valueAfterKSeconds(5, 1000))
}

func valueAfterKSeconds(n int, k int) int {
	mod := int(1e9 + 7) // e is times 10-to-the

	res := make([]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}

	for i := 1; i <= k; i++ {
		for j := 1; j < len(res); j++ {
			res[j] += res[j-1]
			res[j] %= mod
		}
	}

	return res[len(res)-1]
}
