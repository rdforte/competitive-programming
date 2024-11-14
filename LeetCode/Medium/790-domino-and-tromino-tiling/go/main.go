package main

import "fmt"

func main() {
	fmt.Println(numTilings(3) == 5)
	fmt.Println(numTilings(4) == 11)
}

// dynamic programming
// time complexity O(n)

func numTilings(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	var mod int = 1e9 + 7

	k := make([]int, n)
	p := make([]int, n)

	k[0] = 1
	k[1] = 2
	p[1] = 1

	for i := 2; i < n; i++ {
		k[i] = (k[i-1] + k[i-2] + 2*p[i-1]) % mod
		p[i] = (k[i-2] + p[i-1]) % mod
	}

	return k[n-1]
}
