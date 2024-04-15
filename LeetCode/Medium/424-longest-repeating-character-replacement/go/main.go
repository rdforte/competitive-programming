package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AAAAAA", 0))
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	chars := [26]int{}
	maxFrequency := 0
	var maxLen float64

	for l, r := 0, 0; r < len(s); r++ {
		chars[s[r]-'A']++
		maxFrequency = int(math.Max(float64(maxFrequency), float64(chars[s[r]-'A'])))

		for ((r-l)+1)-maxFrequency > k {
			chars[s[l]-'A']--
			l++
		}

		maxLen = math.Max(maxLen, float64((r-l)+1))
	}

	return int(maxLen)
}
