package main

import "fmt"

func main() {
	fmt.Println(numToPower(3, 0), binpow(3, 0))
	fmt.Println(numToPower(3, 1), binpow(3, 1))
	fmt.Println(numToPower(3, 1), binpow(3, 1))
	fmt.Println(numToPower(3, 2), binpow(3, 2))
	fmt.Println(numToPower(3, 3), binpow(3, 3))
	fmt.Println(numToPower(3, 13), binpow(3, 13))
}

// time O(log.n)
func binpow(num, pow int) int {
	res := 1

	for pow > 0 {
		if pow&1 == 1 {
			res *= num
		}
		num *= num
		pow >>= 1
	}

	return res
}

// time: O(n)
func numToPower(num int, pow int) int {
	res := 1
	for i := 1; i <= pow; i++ {
		res *= num
	}

	return res
}
