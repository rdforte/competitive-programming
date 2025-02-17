package main

import "fmt"

func main() {
	h := Constructor()
	h.Hit(1)
	h.Hit(2)
	h.Hit(3)
	fmt.Println(h.GetHits(4))
	h.Hit(300)
	fmt.Println(h.GetHits(300))
	fmt.Println(h.GetHits(301))
	fmt.Println("---------------")
	h2 := ConstructorV2()
	h2.Hit(1)
	h2.Hit(2)
	h2.Hit(3)
	fmt.Println(h2.GetHits(4))
	h2.Hit(300)
	h2.Hit(300)
	fmt.Println(h2.GetHits(300))
	fmt.Println(h2.GetHits(301))
}

type HitCounterV2 struct {
	hitQueue [][]int
	count    int
}

func ConstructorV2() HitCounterV2 {
	return HitCounterV2{
		hitQueue: make([][]int, 0, 300),
	}
}

func (h *HitCounterV2) Hit(timestamp int) {
	h.trimQueue(timestamp)
	if len(h.hitQueue) > 0 && h.hitQueue[len(h.hitQueue)-1][0] == timestamp {
		h.hitQueue[len(h.hitQueue)-1][1]++
	} else {
		h.hitQueue = append(h.hitQueue, []int{timestamp, 1})
	}
	h.count++
}

func (h *HitCounterV2) GetHits(timestamp int) int {
	h.trimQueue(timestamp)
	return h.count
}

func (h *HitCounterV2) trimQueue(timestamp int) {
	trimBy := 0
	for i, time := range h.hitQueue {
		if timestamp-time[0] < 300 {
			break
		}
		h.count -= time[1]
		trimBy = i + 1
	}

	h.hitQueue = h.hitQueue[trimBy:]
}

type HitCounter struct {
	hitQueue []int
}

func Constructor() HitCounter {
	return HitCounter{
		hitQueue: make([]int, 0, 300),
	}
}

func (h *HitCounter) Hit(timestamp int) {
	h.trimQueue(timestamp)
	h.hitQueue = append(h.hitQueue, timestamp)
}

func (h *HitCounter) GetHits(timestamp int) int {
	h.trimQueue(timestamp)
	return len(h.hitQueue)
}

func (h *HitCounter) trimQueue(timestamp int) {
	trimBy := 0
	for i, time := range h.hitQueue {
		if timestamp-time < 300 {
			break
		}
		trimBy = i + 1
	}

	h.hitQueue = h.hitQueue[trimBy:]
}
