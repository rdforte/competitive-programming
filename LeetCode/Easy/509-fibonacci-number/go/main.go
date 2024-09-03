package main

import "fmt"

func main() {
	fmt.Println(0, fib(0))
	fmt.Println(1, fib(1))
	fmt.Println(2, fib(2))
	fmt.Println(3, fib(3))
	fmt.Println(4, fib(4))
	fmt.Println(5, fib(5))
	fmt.Println(6, fib(6))
	fmt.Println(7, fib(7))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	backTwo, backOne := 0, 1
	for i := 2; i <= n; i++ {
		backTwo, backOne = backOne, backTwo+backOne
	}

	return backOne
}
