package main

import "sort"

func main() {

}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	gap := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != gap {
			return false
		}
	}

	return true
}
