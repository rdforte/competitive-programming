package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minimumAddedInteger([]int{4, 20, 16, 12, 8}, []int{14, 18, 10}) == -2)
	fmt.Println(minimumAddedInteger([]int{3, 5, 5, 3}, []int{7, 7}) == 2)
	fmt.Println(minimumAddedInteger([]int{3, 5, 3}, []int{5}) == 0)
	fmt.Println(minimumAddedInteger([]int{9, 4, 3, 9, 4}, []int{7, 8, 8}))
	fmt.Println(minimumAddedInteger([]int{10, 2, 8, 7, 5, 6, 7, 10}, []int{5, 8, 5, 3, 8, 4}))
}

func minimumAddedInteger(n1 []int, n2 []int) int {
	sort.Ints(n1)
	sort.Ints(n2)

	minimum := math.Inf(1)

	for i := 0; i < len(n1); i++ {
		diff := n2[0] - n1[i]
		same := 0
		for n1I, n2I := 0, 0; n1I < len(n1) && n2I < len(n2); n1I++ {
			if n1[n1I]+diff == n2[n2I] {
				same++
				n2I++
			}
		}
		if same >= len(n2) {
			minimum = math.Min(minimum, float64(diff))
		}
	}

	return int(minimum)
}
