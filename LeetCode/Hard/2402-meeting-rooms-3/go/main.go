package main

import (
	"fmt"
	"sort"
)

func main() {
	meetings := [][]int{
		{0, 10},
		{1, 5},
		{2, 7},
		{3, 4},
	}
	fmt.Println(mostBooked(2, meetings))

	meetings = [][]int{
		{1, 20},
		{2, 10},
		{3, 5},
		{4, 9},
		{6, 8},
	}
	fmt.Println(mostBooked(3, meetings))
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	roomsEndTime := make([]int, n)
	roomsCount := make([]int, n)

MEETING:
	for len(meetings) > 0 {
		meetingStart, meetingEnd := meetings[0][0], meetings[0][1]

		roomEndFirstIndex := 0
		for i, endTime := range roomsEndTime {
			if endTime < roomsEndTime[roomEndFirstIndex] {
				roomEndFirstIndex = i
			}

			if meetingStart >= endTime {
				roomsEndTime[i] = meetingEnd
				roomsCount[i]++
				meetings = meetings[1:]
				continue MEETING
			}
		}

		offset := roomsEndTime[roomEndFirstIndex] - meetingStart
		if offset < 0 {
			offset = 0
		}

		roomsEndTime[roomEndFirstIndex] = meetingEnd + offset
		roomsCount[roomEndFirstIndex]++
		meetings = meetings[1:]
	}

	roomNum := 0
	for i, c := range roomsCount {
		if c > roomsCount[roomNum] {
			roomNum = i
		}
	}

	return roomNum
}
