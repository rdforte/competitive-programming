package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getSmallestString("zbbz", 3))
	fmt.Println(getSmallestString("xaxcd", 4))
}

func getSmallestString(s string, k int) string {
	if k == 0 {
		return s
	}

	b := []rune{}
	for _, char := range s {
		if k == 0 {
			b = append(b, char)
			continue
		}

		roundUp := 26 - (char - 'a')
		roundDown := char - 'a'

		if roundUp < roundDown && int(roundUp) <= k {
			b = append(b, 'a')
			k -= int(roundUp)
		} else {
			newChar := char - rune(math.Min(float64(k), float64(char-'a')))
			k -= int(math.Min(float64(k), float64(char-'a')))
			b = append(b, newChar)
		}

	}

	return string(b)
}
