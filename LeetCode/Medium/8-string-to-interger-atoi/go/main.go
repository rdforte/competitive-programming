package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// fmt.Println(myAtoi("  -42"))
	// fmt.Println(myAtoi("1337c0d3"))
	// fmt.Println(myAtoi("0-1"))
	// fmt.Println(myAtoi("words and 987"))
	// fmt.Println(myAtoi("0012"))
	// fmt.Println(myAtoi("0202"))
	// fmt.Println(myAtoi("+-12"))
	fmt.Println(myAtoi("1000522545459"))

	fmt.Println(522545459 > math.MaxInt32)
}

func myAtoi(s string) int {
	sign := 1
	hasSign := false
	num := make([]byte, 0, len(s))

	s = strings.TrimLeft(s, " ")

	for i := range s {
		char := s[i]

		if char == '-' {
			if len(num) > 0 || hasSign {
				break
			}
			sign = -1
			hasSign = true
			continue
		}

		if char == '+' {
			if len(num) > 0 || hasSign {
				break
			}
			hasSign = true
			continue
		}

		if char < '0' || char > '9' {
			break
		}

		num = append(num, char)
	}

	upperBound := math.MaxInt32
	lowerBound := math.MinInt32

	res := 0
	places := 1
	for i := len(num) - 1; i >= 0; i-- {
		n := int(num[i]-'0') * places

		if sign == 1 {
			remainder := upperBound - res
			if n > remainder {
				return upperBound
			}
		} else {
			remainder := int(int(math.Abs(float64(lowerBound + res))))
			if n > remainder {
				return lowerBound
			}
		}

		places *= 10
		res += n
	}

	return res * sign
}
