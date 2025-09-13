package main

import (
	"fmt"
)

func main() {
	fmt.Println(minBitFlips(10, 7))
}

func minBitFlips(start, goal int) int {
	count := 0

	for mask := 1; mask <= max(start, goal); mask <<= 1 {
		startBit := start & mask
		goalBit := goal & mask

		if startBit^goalBit == mask {
			count++
		}
	}

	return count
}
