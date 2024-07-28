package main

import "fmt"

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'0', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '1', '1', '1', '1'},
	}) == 9)
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, 0, len(matrix))
	for range len(matrix) {
		dp = append(dp, make([]int, len(matrix[0])))
	}

	maxLen := 0
	for r := 0; r < len(matrix); r++ {
		dp[r][0] = int(matrix[r][0] - '0')
		maxLen = max(maxLen, dp[r][0])
	}

	for c := 0; c < len(matrix[0]); c++ {
		dp[0][c] = int(matrix[0][c] - '0')
		maxLen = max(maxLen, dp[0][c])
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[r]); c++ {
			if matrix[r][c] == '1' {
				dp[r][c] = min(dp[r][c-1], dp[r-1][c], dp[r-1][c-1]) + 1
				maxLen = max(dp[r][c], maxLen)
			}
		}
	}

	return maxLen * maxLen
}
