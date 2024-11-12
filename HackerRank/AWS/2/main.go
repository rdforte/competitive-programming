package main

import "fmt"

func main() {
	fmt.Println(generateNewArray([]int32{10, 5, 7, 6}, "0101", 2))
	fmt.Println(generateNewArray([]int32{5, 4, 3, 6}, "1100", 5))
}

// TLE

func generateNewArray(arr []int32, state string, m int32) []int32 {
	enabled := make([]bool, len(state))
	for i, v := range state {
		if v == '1' {
			enabled[i] = true
		}
	}
	res := make([]int32, 0, m)
	largest := 0
	for m > 0 {
		zerosRightOne := []int{}
		for i, num := range arr {
			if enabled[i] {
				largest = max(largest, int(num))
			}
			if i > 0 && !enabled[i] && enabled[i-1] {
				zerosRightOne = append(zerosRightOne, i)
			}
		}
		res = append(res, int32(largest))
		for _, i := range zerosRightOne {
			enabled[i] = true
		}
		m--
	}

	return res
}
