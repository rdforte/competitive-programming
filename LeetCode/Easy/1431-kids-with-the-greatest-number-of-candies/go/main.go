package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	greatestCandy := 0
	for _, c := range candies {
		greatestCandy = max(greatestCandy, c)
	}

	res := make([]bool, 0, len(candies))
	for _, c := range candies {
		res = append(res, c+extraCandies >= greatestCandy)
	}

	return res
}
