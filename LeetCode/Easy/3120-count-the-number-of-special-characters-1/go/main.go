package main

import "fmt"

func main() {
	fmt.Println(numberOfSpecialCharacters("aaAbcBC") == 3)
	fmt.Println(numberOfSpecialCharacters("abc") == 0)
	fmt.Println(numberOfSpecialCharacters("abBCab") == 1)
}

func numberOfSpecialCharacters(word string) int {
	lower := [26]int{}
	upper := [26]int{}

	for _, r := range word {
		i := r - 'a'
		if i < 26 && i >= 0 {
			lower[i]++
		} else {
			upper[r-'A']++
		}
	}

	var count int
	for i := 0; i < 26; i++ {
		if lower[i] > 0 && upper[i] > 0 {
			count++
		}
	}

	return count
}
