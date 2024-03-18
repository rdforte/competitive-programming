package main

import "fmt"

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("sadbutsad", "but"))
	fmt.Println(strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {
	l, r, n := 0, 0, 0
	for r < len(haystack) {
		if haystack[r] != needle[n] {
			l++
			r = l
			n = 0
			continue
		}
		if len(needle) == (r-l)+1 {
			return l
		}

		r++
		n++
	}

	return -1
}
