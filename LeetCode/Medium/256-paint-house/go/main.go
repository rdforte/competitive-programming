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
}

// One neat trick that I discovered from a previouse question which involved a cyclic array
// was how to find the index when it passes the end of the array and wraps back around.
// In this question I can not have 2 adjacent houses painted the same colour so in order to find
// the colours of the houses which are not the current I can do.
// (i + 1) % 3 , (i + 2) % 3
//
// I use % 3 because there are always 3 colours.
//
// for example at cost index 1.
// (1 + 1) % 3 = 2, (1 + 2) % 3 = 0

func minCost(costs [][]int) int {
	for i := 1; i < len(costs); i++ {
		for c := 0; c < 3; c++ {
			costs[i][c] = min(costs[i-1][(c+1)%3]+costs[i][c], costs[i-1][(c+2)%3]+costs[i][c])
		}
	}

	nCosts := len(costs)
	return min(costs[nCosts-1][0], costs[nCosts-1][1], costs[nCosts-1][2])
}
