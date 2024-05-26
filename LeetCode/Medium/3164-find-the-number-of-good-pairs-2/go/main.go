package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 4}, []int{1, 3, 4}, 1))
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{2, 4}, 3))
	fmt.Println(numberOfPairs([]int{1, 2, 4, 12}, []int{}, 3))
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	factorCount := make(map[int]int64)

	for _, num := range nums1 {
		for i := 1; i*i <= num; i++ {
			if i*i == num {
				factorCount[i]++
			} else if num%i == 0 {
				factorCount[i]++
				factorCount[num/i]++
			}
		}
	}

	var count int64
	for _, n := range nums2 {
		if c, ok := factorCount[n*k]; ok {
			count += c
		}
	}

	return count
}
