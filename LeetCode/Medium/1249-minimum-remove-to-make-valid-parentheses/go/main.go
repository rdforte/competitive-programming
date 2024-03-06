package main

import "fmt"

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)") == "lee(t(c)o)de")
	fmt.Println(minRemoveToMakeValid("(lee(t(c)o)de)") == "(lee(t(c)o)de)")
	fmt.Println(minRemoveToMakeValid("(((lee(t(c)o)de)") == "(lee(t(c)o)de)")
	fmt.Println(minRemoveToMakeValid(")lee(t(c)o))de)") == "lee(t(c)o)de")
}

func minRemoveToMakeValid(s string) string {
	invalidChars := make([]bool, len(s))
	stack := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}

		if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				invalidChars[i] = true
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		invalidChars[top] = true
	}

	// Note one alternative would be to do res := "" and we can append string(s[i]) but this is extremely inefficient the converstion from byte to string.
	// appending the byte to an array of byte and then doing the conversion once is way more efficient.
	// my runtime went from 1265ms down to 2ms that is a massive performance gain.

	res := []byte{}
	for i, isInvalid := range invalidChars {
		if isInvalid {
			continue
		}
		res = append(res, s[i])
	}

	return string(res)
}
