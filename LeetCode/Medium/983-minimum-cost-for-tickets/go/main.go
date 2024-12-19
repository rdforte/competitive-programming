package main

import "fmt"

func main() {
	fmt.Println("BRUTE FORCE")
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	fmt.Println("DP TOP DOWN")
	fmt.Println(mincostTicketsMemo([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTicketsMemo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}

func mincostTickets(days []int, costs []int) int {
	travellingDay := make([]bool, days[len(days)-1]+1)
	for _, d := range days {
		travellingDay[d] = true
	}

	var dp func(d int) int
	dp = func(d int) int {
		if d >= len(travellingDay) {
			return 0
		}

		if !travellingDay[d] {
			return dp(d + 1)
		}

		return min(dp(d+1)+costs[0], dp(d+7)+costs[1], dp(d+30)+costs[2])
	}

	return dp(0)
}

func mincostTicketsMemo(days []int, costs []int) int {
	travellingDay := make([]bool, days[len(days)-1]+1)

	memo := make([]int, 0, days[len(days)-1]+1)
	for range days[len(days)-1] + 1 {
		memo = append(memo, -1)
	}

	for _, d := range days {
		travellingDay[d] = true
	}

	var dp func(d int) int
	dp = func(d int) int {
		if d >= len(travellingDay) {
			return 0
		}

		if !travellingDay[d] {
			return dp(d + 1)
		}

		if memo[d] == -1 {
			memo[d] = min(dp(d+1)+costs[0], dp(d+7)+costs[1], dp(d+30)+costs[2])
		}

		return memo[d]
	}

	return dp(0)
}
