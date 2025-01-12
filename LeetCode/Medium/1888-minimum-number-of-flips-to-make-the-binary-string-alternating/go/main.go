package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minFlipsBruteForce("111000"))
	fmt.Println(minFlipsBruteForce("010"))
	fmt.Println(minFlipsBruteForce("1110"))

	fmt.Println("OPTIMISED")
	fmt.Println(minFlipsOptimised("111000"))
	fmt.Println(minFlipsOptimised("010"))
	fmt.Println(minFlipsOptimised("1110"))
}

// TIME - the string len is now 2n and we calculate the min on 2 different alternating strings str1 and str2.
// Therefore the time is 2n * 2 which is O(n).
// SPACE - I need to hold the strings in memory which is 2n * 2 plus the string we are comparing = 2n*2 + 2n
// Therefore this gives us O(n)
func minFlipsOptimised(s string) int {
	n := len(s)
	str1, str2 := make([]byte, 0, n*2), make([]byte, 0, n*2)
	for i := range n * 2 {
		if i%2 == 0 {
			str1 = append(str1, '1')
			str2 = append(str2, '0')
		} else {
			str1 = append(str1, '0')
			str2 = append(str2, '1')
		}
	}

	str := []byte(s)
	for i := 0; i < n-1; i++ {
		str = append(str, s[i])
	}

	return min(minFlipsStr(str, str1, n), minFlipsStr(str, str2, n))
}

func minFlipsStr(str, comparison []byte, window int) int {
	res := math.MaxInt64
	totalFlips := 0
	for l, r := 0, 0; r < len(str); r++ {
		if str[r] != comparison[r] {
			totalFlips++
		}
		if (r-l)+1 > window {
			if str[l] != comparison[l] {
				totalFlips--
			}
			l++
		}
		if (r-l)+1 == window {
			res = min(res, totalFlips)
		}
	}

	return res
}

// TLE
// Time = O(n^2)
// Space = O(n)
func minFlipsBruteForce(s string) int {
	n := len(s)
	str1, str2 := make([]byte, 0, n), make([]byte, 0, n)
	for i := range n {
		if i%2 == 0 {
			str1 = append(str1, '1')
			str2 = append(str2, '0')
		} else {
			str1 = append(str1, '0')
			str2 = append(str2, '1')
		}
	}

	minFlips := math.MaxInt64
	for i := 0; i < len(s); i++ {
		str := []byte(s[i:] + s[:i])

		flips1, flips2 := 0, 0
		for i, c := range str {
			if str1[i] != c {
				flips1++
			}
			if str2[i] != c {
				flips2++
			}
		}
		minFlips = min(minFlips, flips1, flips2)
	}

	return minFlips
}
