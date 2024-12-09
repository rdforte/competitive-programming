package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC") == "BANC")
	fmt.Println(minWindow("a", "a") == "a")
	fmt.Println(minWindow("a", "aa") == "")
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd") == "abbbbbcdd")
}

// Space = O(|S| + |T|) where |S| is the length of s and |T| is the length of t.
// Time = O(|S| + |T|) where |S| is the length of s and |T| is the length of t.
func minWindow(s string, t string) string {
	tMp := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		tMp[t[i]]++
	}

	sMp := make(map[byte]int, len(s))

	sMatch, tMatch := 0, len(tMp)

	lp, rp := -1, len(s)

	for l, r := 0, 0; r < len(s); r++ {
		rChar := s[r]

		if val, ok := tMp[rChar]; ok {
			sMp[rChar]++
			if sMp[rChar] == val {
				sMatch++
			}
		}

		for sMatch == tMatch && l <= r {
			if (r - l) < (rp - lp) {
				rp, lp = r, l
			}

			lChar := s[l]
			if val, ok := tMp[lChar]; ok {
				sMp[lChar]--
				if sMp[lChar] < val {
					sMatch--
				}
			}
			l++
		}
	}

	if lp == -1 {
		return ""
	}

	return s[lp : rp+1]
}
