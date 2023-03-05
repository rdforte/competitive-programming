package main

import "fmt"

func lastStoneWeight(stones []int) int {

	var num int = 1
	for i := 1; i <= 3; i++ {
		num = num << i
	}
	return num
}

func main() {
	res := lastStoneWeight(make([]int, 3))
	fmt.Println(res)
}
