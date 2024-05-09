package main

import "fmt"

func main() {
	fmt.Println(partition("baab"))
	fmt.Println(partition("aab"))
}

func partition(s string) [][]string {
	res := [][]string{}
	temp := []string{}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(s) {
			valid := make([]string, len(temp))
			copy(valid, temp)
			res = append(res, valid)
		}

		for j := i; j < len(s); j++ {
			t := s[i : j+1]

			if checkPalindrome(t) {
				temp = append(temp, s[i:j+1])
				dfs(j + 1)
				temp = temp[:len(temp)-1]
			}

		}
	}

	dfs(0)

	return res
}

func checkPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}

	return true
}
