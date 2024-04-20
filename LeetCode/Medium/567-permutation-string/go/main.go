package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "ab") == true)
	fmt.Println("--------------")
	fmt.Println(checkInclusion("ab", "eidbaooo") == true)
	fmt.Println("--------------")
	fmt.Println(checkInclusion("abb", "eidbabooo") == true)
	fmt.Println("--------------")
	fmt.Println(checkInclusion("aaa", "aaaaaaa") == true)
	fmt.Println("--------------")
	fmt.Println(checkInclusion("aaaaaaa", "aaaaaaa") == true)
	fmt.Println("--------------")
	fmt.Println(checkInclusion("hello", "ooolleoooleh") == false)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Chars := [26]int{}
	s2Chars := [26]int{}

	for i, r := range s1 {
		s1Chars[r-'a']++
		s2Chars[s2[i]-'a']++
	}

	var matches int
	for i := 0; i < 26; i++ {
		if s1Chars[i] == s2Chars[i] {
			matches++
		}
	}

	if matches == 26 {
		return true
	}

	for l, r := 1, len(s1); r < len(s2); l, r = l+1, r+1 {
		s2Chars[s2[l-1]-'a']--
		if s2Chars[s2[l-1]-'a'] == s1Chars[s2[l-1]-'a'] {
			matches++
		} else if s2Chars[s2[l-1]-'a']+1 == s1Chars[s2[l-1]-'a'] { // it was equal before but now we have made it not equal
			matches--
		}

		s2Chars[s2[r]-'a']++
		if s2Chars[s2[r]-'a'] == s1Chars[s2[r]-'a'] {
			matches++
		} else if s2Chars[s2[r]-'a']-1 == s1Chars[s2[r]-'a'] { // it was equal before but now we have made it not equal
			matches--
		}

		if matches == 26 {
			return true
		}
	}

	return false
}
