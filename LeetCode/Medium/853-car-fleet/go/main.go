package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{3}, []int{3}))
	fmt.Println(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}

func carFleet(target int, position []int, speed []int) int {
	cars := [][]float64{}
	for i := 0; i < len(position); i++ {
		cars = append(cars, []float64{float64(position[i]), float64(speed[i])})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] < cars[j][0]
	})

	stack := [][]float64{}
	stack = append(stack, cars[len(cars)-1])
	for i := len(cars) - 2; i >= 0; i-- {
		car := cars[i]
		carTimeToTarget := (float64(target) - car[0]) / car[1]
		stackCarTimetoTarget := (float64(target) - stack[len(stack)-1][0]) / stack[len(stack)-1][1]
		if carTimeToTarget > stackCarTimetoTarget {
			stack = append(stack, car)
		}
	}

	return len(stack)
}
