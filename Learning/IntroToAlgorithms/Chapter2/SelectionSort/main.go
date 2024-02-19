package main

import "fmt"

func main() {
	arr := []int{8, 7, 6, 1, 10, 4}
	selectionSort(arr)
	fmt.Println(arr)
}

// time complexity O(n^2)
// space O(1) - acts on original input
func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		smallestIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[smallestIdx] > arr[j] {
				smallestIdx = j
			}
		}
		arr[i], arr[smallestIdx] = arr[smallestIdx], arr[i]
	}
}
