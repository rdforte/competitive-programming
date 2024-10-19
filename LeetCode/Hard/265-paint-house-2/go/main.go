package main

import "fmt"

func main() {
	fmt.Println(minCost([][]int{
		{17, 2, 17},
		{16, 16, 5},
		{14, 3, 19},
	}))
	fmt.Println(minCost([][]int{
		{7, 6, 2},
	}))
	fmt.Println(minCost([][]int{
		{1, 5, 3},
		{2, 9, 4},
	}))
	fmt.Println(minCost([][]int{
		{1, 3},
		{2, 4},
	}))
}

func minCost(costs [][]int) int {
	numColors, numHouses := len(costs[0]), len(costs)

	for i := 1; i < len(costs); i++ {
		for c := 0; c < numColors; c++ {
			minColour := costs[i-1][(c+1)%numColors]
			for prevC := 2; prevC < numColors; prevC++ {
				minColour = min(minColour, costs[i-1][(c+prevC)%numColors])
			}
			costs[i][c] += minColour
		}
	}

	res := costs[numHouses-1][numColors-1]
	for _, c := range costs[numHouses-1] {
		res = min(res, c)
	}

	return res
}
