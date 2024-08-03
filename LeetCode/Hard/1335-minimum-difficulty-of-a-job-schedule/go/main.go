package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("TOP DOWN")
	fmt.Println(minDifficultyTopDown([]int{3, 2, 1}, 2) == 4)
	fmt.Println(minDifficultyTopDown([]int{6, 5, 4, 3, 2, 1}, 2) == 7)
	fmt.Println(minDifficultyTopDown([]int{9, 9, 9}, 4) == -1)
	fmt.Println(minDifficultyTopDown([]int{1, 1, 1}, 3) == 3)
}

func minDifficultyTopDown(jobDifficulty []int, d int) int {
	numJobs := len(jobDifficulty)

	if d > numJobs {
		return -1
	}

	memo := make([][]int, d)

	var dp func(job, day int) int
	dp = func(job, day int) int {
		dayIndex := day - 1
		if len(memo[dayIndex]) == 0 {
			memo[dayIndex] = make([]int, 0, numJobs)
			for range jobDifficulty {
				memo[dayIndex] = append(memo[dayIndex], -1)
			}
		}

		if val := memo[dayIndex][job]; val == -1 {
			lastJobCanDoOnCurDay := numJobs - (d - day)
			maxDifficultyOnDay := 0
			if day == d {
				for i := job; i < lastJobCanDoOnCurDay; i++ {
					maxDifficultyOnDay = max(jobDifficulty[i], maxDifficultyOnDay)
				}
				memo[dayIndex][job] = maxDifficultyOnDay
			} else {
				minDifficulty := math.MaxInt
				for i := job; i < lastJobCanDoOnCurDay; i++ {
					maxDifficultyOnDay = max(jobDifficulty[i], maxDifficultyOnDay)
					minDifficulty = min(minDifficulty, dp(i+1, day+1)+maxDifficultyOnDay)
				}
				memo[dayIndex][job] = minDifficulty
			}
		}

		return memo[dayIndex][job]
	}

	return dp(0, 1)
}
