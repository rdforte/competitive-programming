package main

import "fmt"

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("sadbutsad", "but"))
	fmt.Println(strStr("leetcode", "leeto"))
}

func strStr(haystack string, needle string) int {
	for i, j := 0, 0; j < len(haystack); j++ {
		idx := j - i
		if haystack[j] != needle[idx] {
			j = i
			i++
		}

		if (j-i)+1 == len(needle) {
			return i
		}
	}

	return -1
}
