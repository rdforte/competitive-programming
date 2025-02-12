package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findHighAccessEmployees([][]string{
		{"a", "0549"}, {"b", "0457"}, {"a", "0532"}, {"a", "0621"}, {"b", "0540"},
	}))

	fmt.Println(findHighAccessEmployees([][]string{
		{"cd", "1025"}, {"ab", "1025"}, {"cd", "1046"}, {"cd", "1055"}, {"ab", "1124"}, {"ab", "1120"},
	}))
}

// time = O(u * t * log * t)
// space = O(n)
func findHighAccessEmployees(access_times [][]string) []string {
	userTimes := make(map[string][]int)

	// O(n)
	for _, at := range access_times {
		user := at[0]
		timeStr := at[1]
		hours, _ := strconv.Atoi(timeStr[:2])
		mins, _ := strconv.Atoi(timeStr[2:])
		userTimes[user] = append(userTimes[user], hours*60+mins)
	}

	var res []string

	// u = user
	// t = times
	// O(u * t * log * t)
	for user, times := range userTimes {
		sort.Ints(times) // O(t*log*t)
		for l, r := 0, 0; r < len(times); r++ {
			for r-l > 2 {
				l++
			}
			if times[r]-times[l] < 60 && r-l == 2 {
				res = append(res, user)
				break
			}
		}
	}

	return res
}
