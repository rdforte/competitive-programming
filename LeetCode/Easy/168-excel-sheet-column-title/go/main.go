package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(convertToTitle(701))
}

// Think in Base26.
// Because we start from 1 and there is no concept of 0 this means we have to subtract 1.
// https://neetcode.io/solutions/excel-sheet-column-title
func convertToTitle(columnNumber int) string {
	res := []byte{}

	for columnNumber > 0 {
		rem := (columnNumber - 1) % 26
		char := byte('A' + rem)
		res = append(res, char)
		columnNumber = (columnNumber - 1) / 26
	}

	slices.Reverse(res)
	return string(res)
}
