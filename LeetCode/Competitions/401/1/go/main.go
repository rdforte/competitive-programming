package main

import "fmt"

func main() {
	fmt.Println(numberOfChild(3, 5))
	fmt.Println(numberOfChild(5, 6))
	fmt.Println(numberOfChild(4, 2))
}

func numberOfChild(n int, k int) int {
	n--

	isEven := (k/n)%2 == 0

	var res int
	if isEven {
		res = 0 + (k % n)
	} else {
		res = n - (k % n)
	}

	return res

}
