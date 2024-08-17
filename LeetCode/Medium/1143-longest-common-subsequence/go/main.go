package main

import "fmt"

func main() {
	fmt.Println("TOP DOWN - RECURSIVE")
	fmt.Println(longestCommonSubsequenceTopDown("abcde", "ace") == 3)
	fmt.Println(longestCommonSubsequenceTopDown("abc", "abc") == 3)
	fmt.Println(longestCommonSubsequenceTopDown("abc", "def") == 0)
	fmt.Println("BOTTOM UP - ITERATIVE")
	fmt.Println(longestCommonSubsequenceTopDown("abcde", "ace") == longestCommonSubsequenceBottomUp("abcde", "ace"))
	fmt.Println(longestCommonSubsequenceTopDown("abc", "abc") == longestCommonSubsequenceBottomUp("abc", "abc"))
	fmt.Println(longestCommonSubsequenceTopDown("abc", "def") == longestCommonSubsequenceBottomUp("abc", "def"))
}

// Space and Time complexity is O(t1*t2) for both top down and bottom up
// Note: It was easier to figure out the top down first doing brute force with no memoizaton.
// Once I had the brute force solution I was then able to add memoization.
// When doing the bottom up I can then take the algorithm I had divised in "Top Down" and translate that same logic into an array.
// essentially changing dp(t1+1, t2+1) to dp[t1+1][t2+2].
// It also helped to start towards the end and having the base cases on the outside of the matrix as it was easier to allign
// the indices for both texts.
// We could then iterate our way to the top left corner where this would be our solution.

func longestCommonSubsequenceTopDown(text1 string, text2 string) int {
	memo := make([][]int, len(text1))
	for i := range memo {
		for range text2 {
			memo[i] = append(memo[i], -1)
		}
	}

	var dp func(t1, t2 int) int
	dp = func(t1, t2 int) int {
		if t1 >= len(text1) || t2 >= len(text2) {
			return 0
		}

		if memo[t1][t2] == -1 {
			if text1[t1] == text2[t2] {
				memo[t1][t2] = 1 + dp(t1+1, t2+1) // 1 + dp[t1+1][t2+1] in Bottom Up
			} else {
				memo[t1][t2] = max(dp(t1, t2+1), dp(t1+1, t2))
			}
		}

		return memo[t1][t2]
	}

	return dp(0, 0)
}

func longestCommonSubsequenceBottomUp(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for t1 := len(text1) - 1; t1 >= 0; t1-- {
		for t2 := len(text2) - 1; t2 >= 0; t2-- {
			if text1[t1] == text2[t2] {
				dp[t1][t2] = 1 + dp[t1+1][t2+1] // 1 + dp(t1+1, t2+1) in Top Down.
			} else {
				dp[t1][t2] = max(dp[t1][t2+1], dp[t1+1][t2])
			}
		}
	}

	return dp[0][0]
}
