package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
//
// ] = False
// [ = False
// [] = True
// ()[] = True

func main() {
	fmt.Println(checkIsValidParenthesis("[]") == true)
	fmt.Println(checkIsValidParenthesis("()") == true)
	fmt.Println(checkIsValidParenthesis("{}") == true)
	fmt.Println(checkIsValidParenthesis("[]{}()") == true)
	fmt.Println(checkIsValidParenthesis("[{()}]") == true)
	fmt.Println(checkIsValidParenthesis("(()()[]{}([{[]}]))") == true)
	fmt.Println(checkIsValidParenthesis("]") == false)
	fmt.Println(checkIsValidParenthesis(")") == false)
	fmt.Println(checkIsValidParenthesis("}") == false)
	fmt.Println(checkIsValidParenthesis("{") == false)
	fmt.Println(checkIsValidParenthesis("(") == false)
	fmt.Println(checkIsValidParenthesis("[") == false)
	fmt.Println(checkIsValidParenthesis("{}}") == false)
	fmt.Println(checkIsValidParenthesis("{}]") == false)
	fmt.Println(checkIsValidParenthesis("{})") == false)
}

func checkIsValidParenthesis(str string) bool {
	stack := []byte{}

	for i := 0; i < len(str); i++ {
		bracket := str[i]
		if bracket == '(' || bracket == '{' || bracket == '[' {
			stack = append(stack, bracket)
		} else if len(stack) == 0 {
			return false
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !isCorrespondingCharacter(top, bracket) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isCorrespondingCharacter(char1, char2 byte) bool {
	if char1 == '(' && char2 != ')' {
		return false
	}
	if char1 == '[' && char2 != ']' {
		return false
	}
	if char1 == '{' && char2 != '}' {
		return false
	}

	return true
}
