package main

import (
	"fmt"
)

func main() {
	fmt.Println(compressedString("abcde"))
	fmt.Println(compressedString("a"))
	fmt.Println(compressedString("aaaaaaaaaaaaaabb"))
}

func compressedString(word string) string {
	res := []byte{}

	char := word[0]
	count := '0'

	for _, c := range []byte(word) {
		if c != char || count == '9' {
			res = append(res, byte(count), char)
			count = '1'
			char = c
		} else {
			count++
		}
	}

	res = append(res, byte(count), char)

	return string(res)
}
