package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(alertNames(
		[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
		[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
	))

	fmt.Println(alertNames(
		[]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
		[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
	))

	fmt.Println(alertNames(
		[]string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
		[]string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
	))

	fmt.Println(alertNames(
		[]string{"john", "john", "john"},
		[]string{"23:58", "23:59", "00:01"},
	))
}

func alertNames(keyName []string, keyTime []string) []string {
	var res []string
	times := make(map[string][]int)

	for i, name := range keyName {
		hour, _ := strconv.Atoi(keyTime[i][:2])
		minute, _ := strconv.Atoi(keyTime[i][3:])
		times[name] = append(times[name], (hour*60)+minute)
	}

	for name, t := range times {
		// there is the potential for a later time to be 00:01 so by sorting we create a time loop.
		sort.Ints(t)
		for l, r := 0, 0; r < len(t); r++ {
			for t[r]-t[l] > 60 {
				l++
			}
			if r-l >= 2 {
				res = append(res, name)
				break
			}
		}
	}

	sort.Strings(res)
	return res
}
