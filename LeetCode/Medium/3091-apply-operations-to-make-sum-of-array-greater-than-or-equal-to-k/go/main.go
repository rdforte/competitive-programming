package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minOperations(1) == 0)
	fmt.Println(minOperations(2) == 1)
	fmt.Println(minOperations(11) == 5)
	fmt.Println(minOperations(8) == 4)
}

func minOperations(k int) int {
	sk := math.Ceil(math.Sqrt(float64(k)))
	return (int(sk - 1)) + (int(math.Ceil(float64(k)/sk)) - 1)
}
