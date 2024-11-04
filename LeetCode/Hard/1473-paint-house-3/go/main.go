package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("BRUTE FORCE ////////////////")
	fmt.Println(minCost([]int{0, 0, 0, 0, 0}, [][]int{
		{1, 10},
		{10, 1},
	}, 2, 2, 2))
	fmt.Println(minCost([]int{0, 0, 0, 0, 0}, [][]int{
		{1, 10},
		{10, 1},
		{10, 1},
		{1, 10},
		{5, 1},
	}, 5, 2, 3))
	fmt.Println(minCost([]int{0, 2, 1, 2, 0}, [][]int{
		{1, 10},
		{10, 1},
		{10, 1},
		{1, 10},
		{5, 1},
	}, 5, 2, 3))
	fmt.Println("TOP DOWN")
	fmt.Println(minCostTopDown([]int{0, 0, 0, 0, 0}, [][]int{
		{1, 10},
		{10, 1},
	}, 2, 2, 2))
	fmt.Println(minCostTopDown([]int{0, 0, 0, 0, 0}, [][]int{
		{1, 10},
		{10, 1},
		{10, 1},
		{1, 10},
		{5, 1},
	}, 5, 2, 3))
	fmt.Println(minCostTopDown([]int{0, 2, 1, 2, 0}, [][]int{
		{1, 10},
		{10, 1},
		{10, 1},
		{1, 10},
		{5, 1},
	}, 5, 2, 3))
}

// Brute Force - recursion
func minCost(houses []int, cost [][]int, numHouses, _, targetNeighbourhoods int) int {
	var dp func(house, prevColour, neighbourhood int) int
	dp = func(house, prevColour, neighbourhood int) int {
		if house == numHouses {
			if neighbourhood == targetNeighbourhoods {
				return 0
			} else {
				return -1
			}
		}

		if neighbourhood > targetNeighbourhoods {
			return -1
		}

		minPrice := math.MaxInt

		// already painted
		color := houses[house]
		if color > 0 {
			isSameColour := color == prevColour
			nextHood := neighbourhood
			if !isSameColour {
				nextHood++
			}
			if price := dp(house+1, color, nextHood); price != -1 {
				minPrice = min(minPrice, price)
			}
		} else {
			for i, curPrice := range cost[house] {
				color := i + 1
				isSameColour := prevColour == color
				nextHood := neighbourhood
				if !isSameColour {
					nextHood++
				}

				if price := dp(house+1, color, nextHood); price != -1 {
					minPrice = min(minPrice, curPrice+price)
				}
			}
		}

		if minPrice == math.MaxInt {
			return -1
		}

		return minPrice
	}

	return dp(0, 0, 0)
}

func minCostTopDown(houses []int, cost [][]int, numHouses, numColors, targetNeighbourhoods int) int {
	memo := make([][][]int, 0, numHouses)
	for range numHouses {
		colors := make([][]int, 0, numColors+1)
		for range numColors + 1 {
			hoods := make([]int, 0, targetNeighbourhoods+1)
			for range targetNeighbourhoods + 1 {
				hoods = append(hoods, -2) // -1 is a potential value we can persist so for hoods we have not seen mark as -2
			}
			colors = append(colors, hoods)
		}
		memo = append(memo, colors)
	}

	var dp func(house, prevColour, neighbourhood int) int
	dp = func(house, prevColour, neighbourhood int) int {
		if house == numHouses {
			if neighbourhood == targetNeighbourhoods {
				return 0
			} else {
				return -1
			}
		}

		if neighbourhood > targetNeighbourhoods {
			return -1
		}

		if memo[house][prevColour][neighbourhood] == -2 {
			minPrice := math.MaxInt

			// already painted
			color := houses[house]
			if color > 0 {
				isSameColour := color == prevColour
				nextHood := neighbourhood
				if !isSameColour {
					nextHood++
				}
				if price := dp(house+1, color, nextHood); price != -1 {
					minPrice = min(minPrice, price)
				}
			} else {
				for i, curPrice := range cost[house] {
					color := i + 1
					isSameColour := prevColour == color
					nextHood := neighbourhood
					if !isSameColour {
						nextHood++
					}

					if price := dp(house+1, color, nextHood); price != -1 {
						minPrice = min(minPrice, curPrice+price)
					}
				}
			}

			if minPrice == math.MaxInt {
				memo[house][prevColour][neighbourhood] = -1
			} else {
				memo[house][prevColour][neighbourhood] = minPrice
			}
		}

		return memo[house][prevColour][neighbourhood]
	}

	return dp(0, 0, 0)
}
