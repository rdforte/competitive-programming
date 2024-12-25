package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("IceCreAm") == reverseVowels2Pointers("IceCreAm"))
	fmt.Println(reverseVowels("leetcode") == reverseVowels2Pointers("leetcode"))
}

var mp = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func reverseVowels(s string) string {
	str := []byte(s)
	idxs := make([]int, 0, len(s))

	for i := range s {
		c := s[i]
		if _, ok := mp[c]; ok {
			idxs = append(idxs, i)
		}
	}

	for l, r := 0, len(idxs)-1; l < r; l, r = l+1, r-1 {
		str[idxs[l]], str[idxs[r]] = str[idxs[r]], str[idxs[l]]
	}

	return string(str)
}

// improved requires not storing idx's
func reverseVowels2Pointers(s string) string {
	str := []byte(s)

	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < len(s) && !isVowel(str[l]) {
			l++
		}
		for r >= 0 && !isVowel(str[r]) {
			r--
		}
		if l < r {
			str[l], str[r] = str[r], str[l]
		}
	}

	return string(str)
}

func isVowel(c byte) bool {
	if _, ok := mp[c]; ok {
		return true
	}

	return false
}
