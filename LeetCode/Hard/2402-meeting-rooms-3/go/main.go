package main

import (
	"fmt"
	"math"
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
	// maxMeetingEnd := 0
	sort.Slice(meetings, func(i, j int) bool {
		// maxMeetingEnd = int(math.Max(float64(meetings[i][1]), float64(maxMeetingEnd)))
		// maxMeetingEnd = int(math.Max(float64(meetings[j][1]), float64(maxMeetingEnd)))
		return meetings[i][0] < meetings[j][0]
	})

	roomsEndTime := make([]int, n)
	roomsOccupied := make([]bool, n)
	roomsCount := make([]int, n)

	curTime := 0
TIMES:
	for len(meetings) > 0 {
		for o, roomOccupied := range roomsOccupied {
			if len(meetings) == 0 {
				break TIMES
			}

			if roomOccupied && curTime >= roomsEndTime[o] {
				roomsEndTime[o] = 0
				roomOccupied = false
			}

			if !roomOccupied && meetings[0][0] <= curTime {
				overlap := (curTime - meetings[0][0])
				if overlap < 0 {
					overlap = 0
				}

				roomsEndTime[o] = meetings[0][1] + overlap
				roomsOccupied[o] = true
				roomsCount[o]++
				meetings = meetings[1:]

				curTime = int(math.Max(float64(curTime), float64(roomsEndTime[0])))
			}
		}
	}

	roomNum := 0
	count := roomsCount[0]
	for i, c := range roomsCount {
		if c > count {
			count = c
			roomNum = i
		}
	}

	return roomNum
}
