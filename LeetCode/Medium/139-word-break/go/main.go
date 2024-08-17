package main

import "fmt"

func main() {
	fmt.Println("TOP DOWN")
	fmt.Println(wordBreakTopDown("applepenapple", []string{"apple", "pen"}) == true)
	fmt.Println(wordBreakTopDown("leetcode", []string{"leet", "code"}) == true)
	fmt.Println(wordBreakTopDown("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)

	fmt.Println("BOTTOM UP")
	fmt.Println(wordBreakBottomUp("applepenapple", []string{"apple", "pen"}) == true)
	fmt.Println(wordBreakBottomUp("leetcode", []string{"leet", "code"}) == true)
	fmt.Println(wordBreakBottomUp("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)
}

func wordBreakTopDown(s string, wordDict []string) bool {
	memo := make([]int, 0, len(s))
	for range len(s) {
		memo = append(memo, -1)
	}

	var dp func(i int) bool
	dp = func(i int) bool {
		if i == len(s) {
			return true
		}

		if memo[i] == -1 {
			for _, w := range wordDict {
				end := i + len(w)
				if end > len(s) {
					continue
				}
				word := s[i:end]
				if word == w && dp(end) {
					memo[i] = 1
					break
				}
			}

			if memo[i] == -1 {
				memo[i] = 0
			}
		}

		return memo[i] == 1
	}

	return dp(0)
}

func wordBreakBottomUp(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(dp); i++ {
		for _, w := range wordDict {
			end := i + len(w)
			if end > len(s) {
				continue
			}
			word := s[i:end]
			if word == w && (end-len(w) == 0 || dp[end-len(w)-1]) {
				dp[end-1] = true
			}
		}
	}

	return dp[len(dp)-1]
}
