package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumEnergy([]int{5, 2, -10, -5, 1}, 3) == maximumEnergyDP([]int{5, 2, -10, -5, 1}, 3))
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2) == maximumEnergyDP([]int{-2, -3, -1}, 2))
	fmt.Println(maximumEnergy([]int{5, 2, -10, -6, -2}, 3) == maximumEnergyDP([]int{5, 2, -10, -6, -2}, 3))
	fmt.Println(maximumEnergy([]int{5, -10, 4, 3, 5, -9, 9, -7}, 2) == maximumEnergyDP([]int{5, -10, 4, 3, 5, -9, 9, -7}, 2)) // 23
}

func maximumEnergy(energy []int, k int) int {
	e := make([]float64, len(energy))

	for i := 0; i < k && i < len(energy); i++ {
		e[i] = float64(energy[i])
	}

	for i := k; i < len(energy); i++ {
		e[i] = math.Max(float64(energy[i]), float64(energy[i])+e[i-k])
	}

	maxE := math.Inf(-1)
	for i := len(e) - 1; i >= len(e)-k && i >= 0; i-- {
		maxE = math.Max(e[i], maxE)
	}

	return int(maxE)
}

// using DP we iterate backwards because the max can be at the end
func maximumEnergyDP(energy []int, k int) int {
	dp := make([]int, len(energy))

	maxE := math.Inf(-1)
	for i := len(energy) - 1; i >= len(energy)-k; i-- {
		dp[i] = energy[i]
		maxE = math.Max(float64(dp[i]), maxE)
	}

	for i := len(energy) - 1 - k; i >= 0; i-- {
		dp[i] = int(math.Max(float64(dp[i+k]), float64(energy[i]+dp[i+k])))
		maxE = math.Max(float64(dp[i]), maxE)
	}

	return int(maxE)
}
