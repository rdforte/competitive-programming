package main

import (
	"fmt"
	"reflect"
)

func main() {
	numOne := []int{1, 0, 1}
	numTwo := []int{1, 0, 1}
	fmt.Println(reflect.DeepEqual(addBinaryNumbers(numOne, numTwo), []int{1, 0, 1, 0}))

	numOne = []int{1}
	numTwo = []int{1}
	fmt.Println(reflect.DeepEqual(addBinaryNumbers(numOne, numTwo), []int{1, 0}))

	numOne = []int{0}
	numTwo = []int{1}
	fmt.Println(reflect.DeepEqual(addBinaryNumbers(numOne, numTwo), []int{1}))

	numOne = []int{1, 0}
	numTwo = []int{0, 1}
	fmt.Println(reflect.DeepEqual(addBinaryNumbers(numOne, numTwo), []int{1, 1}))

	numOne = []int{0}
	numTwo = []int{0}
	fmt.Println(reflect.DeepEqual(addBinaryNumbers(numOne, numTwo), []int{0}))

	numOne = []int{1, 0, 0}
	numTwo = []int{0, 0, 1}
	fmt.Println(reflect.DeepEqual(addBinaryNumbers(numOne, numTwo), []int{1, 0, 1}))
}

// assume len of a == b
func addBinaryNumbers(a []int, b []int) []int {
	res := make([]int, len(a))

	carry := 0
	for i := len(res) - 1; i >= 0; i-- {
		sum := a[i] + b[i] + carry
		if sum >= 2 {
			res[i] = 0
			carry = 1
		} else if sum == 1 {
			carry = 0
			res[i] = 1
		} else {
			carry = 0
			res[i] = 0
		}
	}

	if carry == 1 {
		res = append([]int{1}, res...)
	}

	return res
}
