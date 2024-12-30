package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(answerString("dbca", 2) == "dbc")
	fmt.Println(answerString("gggg", 4) == "g")
	fmt.Println(answerString("bif", 2) == "if")
	fmt.Println(answerString("aann", 2) == "nn")
}

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	friendWordLen := len(word) - (numFriends - 1)

	largest := ""
	for i := friendWordLen; i <= len(word); i++ {
		for j := i - friendWordLen; j < i; j++ {
			if strings.Compare(largest, word[j:i]) == -1 {
				largest = word[j:i]
			}
		}
	}

	return largest
}
