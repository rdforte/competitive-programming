package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
}

func maxVowels(s string, k int) int {
	res, curVowels := 0, 0
	for l, r := 0, 0; r < len(s); r++ {
		if isVowel(s[r]) {
			curVowels++
		}
		if (r-l)+1 > k {
			if isVowel(s[l]) {
				curVowels--
			}
			l++
		}
		res = max(res, curVowels)
	}

	return res
}

func isVowel(char byte) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
