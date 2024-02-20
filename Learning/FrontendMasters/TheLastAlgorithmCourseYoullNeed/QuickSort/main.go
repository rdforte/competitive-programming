package main

import "fmt"

// Quick sort best case time complexity is O(nlogn) and worst case time complexity is O(n^2)
// Quick sort is usually picked over merge sort because it is faster on average due to merge sorts large constant factor.
// Because our quick sort algorithm does the sorting in place our memory is O(1) as we do not create any new memory besides the constants.
// Quick Sort is reffered to as an in-place sorting algorithm

func main() {
	arr := []int{1, 5, 3, 2, 0}
	quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) {
	var qs func(lo, hi int)
	qs = func(lo, hi int) {
		if lo >= hi {
			return
		}

		pivotIdx := partition(arr, lo, hi)

		qs(lo, pivotIdx-1)
		qs(pivotIdx+1, hi)
	}

	qs(0, len(arr)-1)
}

// We pick the last element here. A slight optimisation would be to pick the middle because if the array was in revers our
// quick sort would be O(n^2)
func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[hi], arr[idx] = arr[idx], arr[hi]

	return idx
}
