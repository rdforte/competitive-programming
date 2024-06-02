package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(countDays(10, [][]int{
		{5, 7},
		{1, 3},
		{9, 10},
	}))

	fmt.Println(countDays(5, [][]int{
		{2, 4},
		{1, 3},
	}))

	fmt.Println(countDays(6, [][]int{
		{1, 6},
	}))

	fmt.Println(countDays(6, [][]int{
		{3, 3},
	}))
	fmt.Println(countDays(10, [][]int{
		{2, 3},
		{5, 6},
	}))
}

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	meets := [][]int{}
	meets = append(meets, meetings[0])

	i := 0
	for _, m := range meetings[1:] {
		if m[0] >= meets[i][0] && m[0] <= meets[i][1] {
			meets[i][1] = int(math.Max(float64(meets[i][1]), float64(m[1])))
		} else {
			meets = append(meets, m)
			i++
		}
	}

	daysBetween := 0
	for i := 1; i < len(meets); i++ {
		diff := (meets[i][0] - meets[i-1][1]) - 1
		if diff > 0 {
			daysBetween += diff
		}
	}

	daysBefore := meets[0][0] - 1
	daysAfter := days - meets[len(meets)-1][1]

	return daysBefore + daysBetween + daysAfter
}
