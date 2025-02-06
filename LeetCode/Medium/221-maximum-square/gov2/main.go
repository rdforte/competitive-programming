package main

import "fmt"

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		{'0', '1'}, {'1', '0'},
	}))
	fmt.Println(maximalSquare([][]byte{
		{'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'0', '0', '1', '1', '1'},
	}))
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, 0, len(matrix))
	for range len(matrix) {
		m := make([]int, len(matrix[0]))
		dp = append(dp, m)
	}

	maxSquare := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == '0' {
				continue
			}
			dp[r][c] = 1

			if c-1 >= 0 && r-1 >= 0 {
				dp[r][c] = min(dp[r-1][c-1]+dp[r][c], dp[r][c-1]+dp[r][c], dp[r-1][c]+dp[r][c])
			}
			maxSquare = max(maxSquare, dp[r][c])
		}
	}

	return maxSquare * maxSquare
}
