package main

import "fmt"

func main() {
	arr := []int{5, 1, 3, 0}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(arr, lo, mid)
	mergeSort(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

func merge(arr []int, lo, mid, hi int) {
	leftArrLen := mid - lo + 1
	rightArrLen := hi - mid

	left := make([]int, leftArrLen)
	right := make([]int, rightArrLen)

	for i := 0; i < leftArrLen; i++ {
		left[i] = arr[lo+i]
	}

	for i := 0; i < rightArrLen; i++ {
		right[i] = arr[mid+i+1]
	}

	l := 0
	r := 0
	k := lo

	for l < leftArrLen && r < rightArrLen {
		if left[l] < right[r] {
			arr[k] = left[l]
			l++
		} else {
			arr[k] = right[r]
			r++
		}
		k++
	}

	for l < leftArrLen {
		arr[k] = left[l]
		l++
		k++
	}

	for r < rightArrLen {
		arr[k] = right[r]
		r++
		k++
	}
}
