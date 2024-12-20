package main

import "fmt"

func main() {
	fmt.Println("Brute Force ---------------")
	fmt.Println(IsInterleave("ab", "cd", "cabd") == true)
	fmt.Println(IsInterleave("aabcc", "dbbca", "aadbbcbcac") == true)
	fmt.Println(IsInterleave("aabcc", "dbbca", "aadbbbaccc") == false)
	fmt.Println(IsInterleave("", "", "") == true)
	fmt.Println("DP Memo ---------------")
	fmt.Println(IsInterleaveMemo("ab", "cd", "cabd") == true)
	fmt.Println(IsInterleaveMemo("aabcc", "dbbca", "aadbbcbcac") == true)
	fmt.Println(IsInterleaveMemo("aabcc", "dbbca", "aadbbbaccc") == false)
	fmt.Println(IsInterleaveMemo("", "", "") == true)
	fmt.Println(IsInterleaveMemo("a", "b", "a") == false)
}

func IsInterleave(s1, s2, s3 string) bool {
	var dp func(i, j, k int) bool
	dp = func(i, j, k int) bool {
		if k == len(s3) {
			return true
		}

		isInterleaving := false
		if i < len(s1) && s1[i] == s3[k] && dp(i+1, j, k+1) {
			isInterleaving = true
		}

		if j < len(s2) && s2[j] == s3[k] && dp(i, j+1, k+1) {
			isInterleaving = true
		}

		return isInterleaving
	}

	return dp(0, 0, 0)
}

func IsInterleaveMemo(s1, s2, s3 string) bool {
	memo := make([][][]int, 0, len(s1)+1)
	for range len(s1) + 1 {
		s2Memo := make([][]int, 0, len(s2)+1)
		for range len(s2) + 1 {
			s3Memo := make([]int, 0, len(s3))
			for range len(s3) {
				s3Memo = append(s3Memo, -1)
			}
			s2Memo = append(s2Memo, s3Memo)
		}
		memo = append(memo, s2Memo)
	}

	var dp func(i, j, k int) int
	dp = func(i, j, k int) int {
		if k == len(s3) {
			if i >= len(s1) && j >= len(s2) {
				return 1
			}
			return 0
		}

		if memo[i][j][k] == -1 {
			iMatch := i < len(s1) && s1[i] == s3[k] && dp(i+1, j, k+1) == 1
			jMatch := j < len(s2) && s2[j] == s3[k] && dp(i, j+1, k+1) == 1
			if iMatch || jMatch {
				memo[i][j][k] = 1
			} else {
				memo[i][j][k] = 0
			}
		}

		return memo[i][j][k]
	}

	if dp(0, 0, 0) == 1 {
		return true
	}

	return false
}
