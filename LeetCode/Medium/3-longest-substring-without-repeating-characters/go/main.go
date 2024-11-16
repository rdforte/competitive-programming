package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int, len(s))
	maxLen := 0

	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := seen[s[r]]; ok {
			prevL := l
			l = seen[s[r]] + 1
			for prevL < l {
				delete(seen, s[prevL])
				prevL++
			}
		}
		maxLen = max(maxLen, (r-l)+1)
		seen[s[r]] = r
	}

	return maxLen
}
