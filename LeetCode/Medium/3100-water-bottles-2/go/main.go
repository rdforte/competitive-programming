package main

import "fmt"

func main() {
	fmt.Println(maxBottlesDrunk(13, 6))
	fmt.Println(maxBottlesDrunk(10, 3))
}

// Question 2
func maxBottlesDrunk(fullBottles int, numExchange int) int {
	bottlesDrunk := fullBottles
	emptyBottles := fullBottles
	fullBottles = 0

	for emptyBottles >= numExchange || fullBottles > 0 {
		if fullBottles > 0 && emptyBottles < numExchange {
			bottlesDrunk += fullBottles
			emptyBottles += fullBottles
			fullBottles = 0
			continue
		}
		fullBottles++
		emptyBottles -= numExchange
		numExchange++
	}

	return bottlesDrunk
}
