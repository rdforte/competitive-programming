package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumLengthSubstring("bcbbbcba") == 4)
	fmt.Println(maximumLengthSubstring("aaaa") == 2)
	fmt.Println(maximumLengthSubstring("aadaad") == 4)
}

func maximumLengthSubstring(s string) int {
	chars := make([][]int, 26)
	var maxLen float64
	l, r := 0, 0

	for r < len(s) {
		c := s[r] - 'a'
		chars[c] = append(chars[c], r)
		if len(chars[c]) > 2 {
			for len(chars[c]) > 2 {
				chars[s[l]-'a'] = chars[s[l]-'a'][1:]
				l++
			}
		}
		maxLen = math.Max(maxLen, float64((r-l)+1))
		r++
	}

	return int(maxLen)
}
