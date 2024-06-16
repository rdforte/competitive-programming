package main

import "fmt"

// A nil slice is an empty slice so to assert if a slice is empty we should always check the length
func main() {
	var a []int            // nil slice
	b := make([]int, 0, 0) // empty slice

	fmt.Println(len(a) == 0)
	fmt.Println(len(b) == 0)
}
