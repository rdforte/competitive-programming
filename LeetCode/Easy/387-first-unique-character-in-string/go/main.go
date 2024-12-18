package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode") == firstUniqueCharSimple("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode") == firstUniqueCharSimple("loveleetcode"))
	fmt.Println(firstUniqChar("aabb") == firstUniqueCharSimple("aabb"))
}

func firstUniqChar(s string) int {
	chars := make([]int, 26)
	queue := make([]int, 0, len(s))

	for i := range s {
		c := s[i]

		chars[c-'a']++
		queue = append(queue, i)

		for len(queue) > 0 && chars[s[queue[0]]-'a'] > 1 {
			queue = queue[1:]
		}
	}

	if len(queue) > 0 {
		return queue[0]
	}

	return -1
}

func firstUniqueCharSimple(s string) int {
	chars := make([]int, 26)

	for i := range s {
		c := s[i]
		chars[c-'a']++
	}

	for i := range s {
		if chars[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
