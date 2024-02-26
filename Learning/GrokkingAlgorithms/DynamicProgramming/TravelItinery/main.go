package main

import (
	"fmt"
	"math"
)

func main() {
	itineries := []Itinery{
		{0.5, 7},
		{0.5, 6},
		{1, 9},
		{2, 9},
		{0.5, 8},
	}
	fmt.Println(optimiseTravelItinery(itineries, 0.5, 2))
}

func optimiseTravelItinery(itineries []Itinery, interval, maxTime float64) int {
	dp := make([][]float64, 0, len(itineries)+1)
	for i := 0; i < len(itineries)+1; i++ {
		dp = append(dp, make([]float64, int(maxTime/interval)+1))
	}

	for r := 1; r < len(dp); r++ {
		itinery := itineries[r-1]
		for c := 0; c < len(dp[r]); c++ {
			if itinery.time > float64(c)*interval {
				dp[r][c] = dp[r-1][c]
			} else {
				prevIndex := math.Max((float64(c)*interval-itinery.time)/interval, 0)
				dp[r][c] = math.Max(dp[r-1][c], itinery.rating+dp[r-1][int(prevIndex)])
			}
		}
	}

	return int(dp[len(dp)-1][len(dp[len(dp)-1])-1])
}

type Itinery struct {
	time   float64
	rating float64
}
