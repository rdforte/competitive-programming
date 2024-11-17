package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("1337c0d3"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("0012"))
	fmt.Println(myAtoi("0202"))
	fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("1000000000000000000000000000000522545459"))

	fmt.Println(math.MaxInt32)
}

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")

	if len(s) == 0 {
		return 0
	}

	sign := 1

	total := 0
	// figure out sign
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for i := range s {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			break
		}

		if (total > math.MaxInt32/10) || (total == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		total = (total * 10) + digit
	}

	return total * sign
}
