package main

import "fmt"

func main() {
	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))
	fmt.Println(sumOfTheDigitsOfHarshadNumber(23))
}

// Question 1
func sumOfTheDigitsOfHarshadNumber(x int) int {
	sum := 0

	num := x
	for num > 0 {
		n := num % 10
		num = num / 10
		sum += n
	}

	if x%sum == 0 {
		return sum
	}

	return -1
}
