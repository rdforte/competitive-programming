package main

import (
	"fmt"
	"strings"
)

func main() {
	res := fullJustify([]string{
		"This",
		"is",
		"an",
		"example",
		"of",
		"text",
		"justification.",
	}, 16)
	for _, r := range res {
		fmt.Printf("=%s=\n", r)
	}

	fmt.Println()
	fmt.Println()

	res = fullJustify([]string{
		"what",
		"must",
		"be",
		"acknowledgment",
		"shall",
		"be",
	}, 16)
	for _, r := range res {
		fmt.Printf("=%s=\n", r)
	}
}

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var curWords []string
	lineLen := 0

	for _, w := range words {
		if len(w)+lineLen > maxWidth {
			curWords[len(curWords)-1] = strings.TrimSuffix(curWords[len(curWords)-1], " ")
			lineLen--

			spacesToAdd := maxWidth - lineLen

			str := ""
			if len(curWords) == 1 {
				str = curWords[0] + strings.Repeat(" ", spacesToAdd)
			} else {
				spacesPerWord := spacesToAdd / (len(curWords) - 1)
				spacesRemaining := spacesToAdd % (len(curWords) - 1)

				for i := 0; i < len(curWords)-1; i++ {
					curWords[i] = curWords[i] + strings.Repeat(" ", spacesPerWord)
				}

				for i := range spacesRemaining {
					curWords[i] = curWords[i] + " "
				}

				str = strings.Join(curWords, "")
			}

			res = append(res, str)
			lineLen = 0
			curWords = nil
		}

		curWords = append(curWords, w+" ")
		lineLen += len(w) + 1
	}

	// last line
	curWords[len(curWords)-1] = strings.TrimSuffix(curWords[len(curWords)-1], " ")
	lineLen--
	spacesToAdd := maxWidth - lineLen
	curWords[len(curWords)-1] = curWords[len(curWords)-1] + strings.Repeat(" ", spacesToAdd)
	str := strings.Join(curWords, "")
	res = append(res, str)

	return res
}
