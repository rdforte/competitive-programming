package main

import "fmt"

func passThePillow(n int, time int) int {
	t := (time % ((n - 1) * 2)) + 1
	if t <= n {
		return t
	}
	return n - (t - n)
}

func main() {
	fmt.Printf("%v", passThePillow(4, 5))
}
