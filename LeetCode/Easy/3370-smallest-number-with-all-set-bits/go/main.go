package main

import "fmt"

func main() {
	fmt.Println(smallestNumber(5) == smallestNumberV2(5))
	fmt.Println(smallestNumber(10) == smallestNumberV2(10))
	fmt.Println(smallestNumber(3) == smallestNumberV2(3))
	fmt.Println(smallestNumber(1) == smallestNumberV2(1))
}

func smallestNumber(n int) int {
	bitNums := 0
	for n > 0 {
		bitNums++
		n /= 2
	}

	res := 0

	for i := 0; i < bitNums; i++ {
		res += (1 << i)
	}

	return res
}

func smallestNumberV2(n int) int {
	x := 1
	for x < n {
		x = (x << 1) | 1 // shift left and use OR bitwise operator.
	}
	return x
}
