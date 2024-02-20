package main

import "fmt"

func main() {
	fmt.Println(linearSearch[float32]([]float32{1, 2, 3}, 2) == true)
	fmt.Println(linearSearch[float32]([]float32{1, 2, 3}, 5) == false)
}

func linearSearch[T comparable](s []T, n T) bool {
	for i := range s {
		if s[i] == n {
			return true
		}
	}
	return false
}
