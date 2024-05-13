package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findePermutationDifference("abc", "bac"))
	fmt.Println(findePermutationDifference("abcde", "edbac"))
}

func findePermutationDifference(s string, t string) int {
	si := [26]int{}
	ti := [26]int{}

	for i := 0; i < len(s); i++ {
		si[s[i]-'a'] = i
		ti[t[i]-'a'] = i
	}

	diff := 0
	for i := 0; i < 26; i++ {
		diff += int(math.Abs(float64(si[i] - ti[i])))
	}

	return diff
}
