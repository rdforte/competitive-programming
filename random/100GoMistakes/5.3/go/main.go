package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"
	fmt.Println(strings.TrimSuffix(s, "lo")) // hel - exact match trim once
	fmt.Println(strings.TrimSuffix(s, "ol")) // hello - exact match trim once

	fmt.Println("---")

	fmt.Println(strings.TrimPrefix(s, "he")) // llo - exact match trim once
	fmt.Println(strings.TrimPrefix(s, "eh")) // hello - exact match trim once

	fmt.Println("---")

	fmt.Println(strings.TrimRight(s, "ol"))  // he - trims all characters in character set up until character not in set.
	fmt.Println(strings.TrimRight(s, "ole")) // h - trims all characters in character set up until character not in set.

	fmt.Println("---")

	fmt.Println(strings.TrimLeft(s, "eh"))  // llo - trims all characters in character set up until character not in set.
	fmt.Println(strings.TrimLeft(s, "leh")) // o - trims all characters in character set up until character not in set.
}
