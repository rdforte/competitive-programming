package main

import "fmt"

// note arrays do not have a capacity where a slice does

func main() {
	a := [5]int{1, 2, 3, 4, 5} // has a fixed length and no capacity.
	a[0] = 6
	fmt.Println(a)
	// a = append(a, 6) // this will not work because a is an array and not a slice

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // has a length and capacity

	newS := make([]int, 0, len(s)) // makes an array with len 0 and capacity set to length of arr.

	for _, num := range s {
		newS = append(newS, num) // because len is 0 we can use append. but our array is initialised with the capacity it needs.
	}

	fmt.Println(newS)
}
