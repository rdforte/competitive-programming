package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(isValid("12") == false)
	fmt.Println(isValid("Hello") == true)
	fmt.Println(isValid("hello") == true)
	fmt.Println(isValid("he99") == true)
	fmt.Println(isValid("AbZ") == true)
	fmt.Println(isValid("AbZ%") == false)
	fmt.Println(isValid("a3$e") == false)
	fmt.Println(isValid("UuE6") == false)
}

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	hasChars, _ := regexp.MatchString(`[a-zA-Z0-9]+`, word)
	hasSpecialChars, _ := regexp.MatchString(`[^a-zA-Z0-9]+`, word)
	hasVowel, _ := regexp.MatchString(`[aeiouAEIOU]+`, word)
	hasConsonent, _ := regexp.MatchString(`[^aeiouAEIOU0-9]+`, word)

	return hasChars && hasVowel && hasConsonent && !hasSpecialChars
}
