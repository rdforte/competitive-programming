package main

import (
	"fmt"
	"time"
)

func main() {
	year := 2018
	month := time.January
	day := 1
	hour := 2
	minute := 3
	second := 10
	nsec := 5

	window := time.Second * 60
	t := time.Date(year, month, day, hour, minute, second, nsec, time.UTC)

	r := New(window, 3)
	fmt.Println(r.IsLimit())
	r.AddRequest(t)
	fmt.Println(r.IsLimit())
	r.AddRequest(time.Date(year, month, day, hour, minute, 59, nsec, time.UTC))
	fmt.Println(r.IsLimit())
	r.AddRequest(time.Date(year, month, day, hour, minute, 59, 10, time.UTC))
	fmt.Println(r.IsLimit())
	r.AddRequest(time.Date(year, month, day, hour, minute, 59, 10, time.UTC))
	fmt.Println(r.IsLimit()) // rate limited
	r.AddRequest(time.Date(year, month, day, hour, minute+1, 59, nsec, time.UTC))
	fmt.Println(r.IsLimit()) // bucket reset to 0
}

// Fixed Window Counter is a simple algorithm.
// We just keep track of the request count for the current time window.
// Downside is that an increase in traffic at the bounds of our window does not cause the
// limit to be hit as once the requests come into a new time window we do not cater for the
// requests in the previouse window.

func New(windowSize time.Duration, capacity int) FWC {
	// Add 1 to the capacity because if capacity is say 10/10 then they should not be at the limit.
	// Its only when they exceed the limit that they are at capacity and should be rate limited.
	return FWC{windowSize: windowSize, cap: capacity + 1}
}

type FWC struct {
	windowSize   time.Duration
	startTime    time.Time
	requestCount int
	cap          int
}

func (f *FWC) AddRequest(time time.Time) {
	bucketTime := time.Truncate(f.windowSize)
	if bucketTime.After(f.startTime) {
		f.startTime = bucketTime
		f.requestCount++
		return
	}

	f.requestCount = min(f.cap, f.requestCount+1)
}

func (f *FWC) IsLimit() bool {
	return f.requestCount == f.cap
}
