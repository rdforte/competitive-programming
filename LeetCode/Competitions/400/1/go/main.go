package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumChairs("EEEEEEE"))
	fmt.Println(minimumChairs("ELELEEL"))
	fmt.Println(minimumChairs("ELEELEELLL"))
}

func minimumChairs(s string) int {
	count := 0.0
	maxChairs := 0.0
	for _, c := range s {
		if c == 'E' {
			count++
		} else {
			count--
		}
		maxChairs = math.Max(count, maxChairs)
	}

	return int(maxChairs)
}
