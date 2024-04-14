package main

import "fmt"

func main() {
	fmt.Println(findLatestTime("11:54"))
	fmt.Println(findLatestTime("1?:?4"))
	fmt.Println(findLatestTime("0?:5?"))
	fmt.Println(findLatestTime("??:3?"))
	fmt.Println(findLatestTime("??:??"))
	fmt.Println(findLatestTime("?1:?6"))
}

func findLatestTime(s string) string {
	newTime := ""

	if s[0] == '?' {
		if s[1] == '?' || s[1] == '0' || s[1] == '1' {
			newTime += "1"
		} else {
			newTime += "0"
		}
	} else {
		newTime += string(s[0])
	}

	if s[1] == '?' {
		if newTime[0] == '0' {
			newTime += "9"
		} else {
			newTime += "1"
		}
	} else {
		newTime += string(s[1])
	}

	newTime += ":"

	if s[3] == '?' {
		newTime += "5"
	} else {
		newTime += string(s[3])
	}

	if s[4] == '?' {
		newTime += "9"
	} else {
		newTime += string(s[4])
	}

	return newTime
}

