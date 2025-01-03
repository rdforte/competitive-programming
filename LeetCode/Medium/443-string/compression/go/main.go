package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'}
	fmt.Println(compress(b))
	fmt.Println(string(b))
	fmt.Println("---------------------------------")

	b = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(b))
	fmt.Println(string(b))
	fmt.Println("---------------------------------")

	b = []byte{'a'}
	fmt.Println(compress(b))
	fmt.Println(string(b))
	fmt.Println("---------------------------------")

	b = []byte{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'a'}
	fmt.Println(compress(b))
	fmt.Println(string(b))
	fmt.Println("---------------------------------")

	b = []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}
	fmt.Println(compress(b))
	fmt.Println(string(b))
	fmt.Println("---------------------------------")
}

func compress(chars []byte) int {
	res := 0
	for i := 0; i < len(chars); {
		groupLen := 1
		for j := i + groupLen; j < len(chars) && chars[i+groupLen] == chars[i]; j++ {
			groupLen++
		}

		chars[res] = chars[i]

		if groupLen > 1 {
			str := strconv.Itoa(groupLen)
			for j := 0; j < len(str); j++ {
				c := str[j]
				chars[res+1+j] = c
			}
			res += len(str)
		}
		res += 1
		i += groupLen
	}

	return res
}
