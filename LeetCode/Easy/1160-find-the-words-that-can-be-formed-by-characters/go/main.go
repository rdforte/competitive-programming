package main

import "fmt"

func main() {
	fmt.Println(countCharactersV2([]string{
		"cat",
		"bt",
		"hat",
		"tree",
	}, "atach"))
}

func countCharacters(words []string, chars string) int {
	// O(chars)
	charCount := make(map[byte]int)
	for i := range chars {
		c := chars[i]
		charCount[c]++
	}

	res := 0

	// O(words * chars)
OUTER:
	for _, w := range words {
		wordMp := make(map[byte]int)
		for i := range w {
			c := w[i]
			wordMp[c]++
		}

		for k, val := range wordMp {
			if v, ok := charCount[k]; !ok || val > v {
				continue OUTER
			}
		}

		res += len(w)
	}

	return res
}

func countCharactersV2(words []string, chars string) int {
	charCount := make([]int, 26)
	// O(chars)
	for i := range chars {
		c := chars[i]
		charCount[c-'a']++
	}

	res := 0

	// O(words * chars)
OUTER:
	for _, w := range words {
		wordMp := make([]int, 26)
		for i := range w {
			c := w[i]
			wordMp[c-'a']++
		}

		for i, count := range wordMp {
			if count > charCount[i] {
				continue OUTER
			}
		}

		res += len(w)
	}

	return res
}
