package main

import "fmt"

func main() {
	arr := []int{7, 6, 5, 4, 1, 9, 2}
	bubbleSort(arr)
	fmt.Println(arr)
}

// O(n^2) time
// Space O(1) as acts on input and doesn't create new memory
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
