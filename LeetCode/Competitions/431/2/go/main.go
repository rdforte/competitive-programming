package main

import "fmt"

func main() {
	fmt.Println(calculateScore("aczzx"))
	fmt.Println(calculateScore("abcdef"))
	fmt.Println(calculateScore("zadavyayobbgqsexaabk"))
}

func calculateScore(s string) int64 {
	chars := make([][]int, 26)

	var res int64
	for i, c := range s {
		need := 'z' - (c - 'a')

		if len(chars[need-'a']) > 0 {
			top := chars[need-'a'][len(chars[need-'a'])-1]
			res += int64(i - top)
			chars[need-'a'] = chars[need-'a'][:len(chars[need-'a'])-1]
		} else {
			chars[c-'a'] = append(chars[c-'a'], i)
		}
	}

	return res
}
