package main

import (
	"fmt"
)

func main() {
	fmt.Println(minAnagramLength("abba"))
	fmt.Println(minAnagramLength("cdef"))
	fmt.Println(minAnagramLength("abccccaaabccccaaabccccaa"))
	fmt.Println(minAnagramLength("ababc"))
	fmt.Println(minAnagramLength("z"))
}

func minAnagramLength(s string) int {
	chars := [26]int{}

	for _, c := range s {
		chars[c-'a']++
	}

	div := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] > 0 {
			div = gcd(div, chars[i])
		}
	}

	return len(s) / div
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
