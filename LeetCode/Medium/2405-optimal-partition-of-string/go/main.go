package main

import "fmt"

func main() {
	fmt.Println(partitionString("abacaba"))
	fmt.Println(partitionString("ssssss"))
}

// space complexity = O(26) = O(1).
// time complexity = O(len(s)*26) = O(s) therefore O(n) time.
func partitionString(s string) int {
	chars := make([]bool, 26)
	chars[s[0]-'a'] = true

	count := 1

	for i := 1; i < len(s); i++ {
		char := s[i]

		if chars[char-'a'] {
			for c := range chars {
				chars[c] = false
			}
			count++
		}
		chars[char-'a'] = true
	}

	return count
}
