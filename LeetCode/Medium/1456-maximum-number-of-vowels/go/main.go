package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
}

func maxVowels(s string, k int) int {
	maxVowels := 0
	currentVowels := 0
	for l, r := 0, 0; r < len(s); r++ {
		if r-l >= k {
			if isVowel(s[l]) {
				currentVowels--
			}
			l++
		}

		if isVowel(s[r]) {
			currentVowels++
		}

		maxVowels = int(math.Max(float64(maxVowels), float64(currentVowels)))
	}

	return maxVowels
}

func isVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
