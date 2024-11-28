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

	r := New(window, 2)
	fmt.Println(r.AddRequest(t.Add(time.Second * 10)))
	fmt.Println(r.AddRequest(t.Add(time.Second * 10)))
	fmt.Println(r.AddRequest(t.Add(time.Minute * 1))) // new bucket
	fmt.Println(r.AddRequest(t.Add(time.Minute * 2))) // new bucket
	fmt.Println(r.AddRequest(t.Add(time.Minute * 2))) // new bucket
	fmt.Println(r.AddRequest(t.Add(time.Minute * 2))) // new bucket
	fmt.Println(r.AddRequest(t.Add(time.Minute * 4))) // new bucket
}

func New(interval time.Duration, capacity int) SWC {
	return SWC{interval: interval, capacity: capacity + 1}
}

type SWC struct {
	interval time.Duration
	b1       Bucket
	b2       Bucket
	capacity int
}

type Bucket struct {
	startTime time.Time
	count     int
}

func (s *SWC) AddRequest(time time.Time) bool {
	curTime := time.Truncate(s.interval)

	// Add to existing bucket
	if curTime.Equal(s.b2.startTime) {
		s.b2.count = min(s.capacity, s.b2.count+1)
		return s.isLimit()
	}

	if curTime.After(s.b2.startTime) {
		// shift buckets
		s.b1 = s.b2
		// set time on new bucket
		s.b2.startTime = curTime
		s.b2.count = 1
	}

	if curTime.Add(-s.interval).After(s.b1.startTime) {
		s.b1.startTime = curTime.Add(-s.interval)
		s.b1.count = 0
	}

	return s.isLimit()
}

func (s *SWC) isLimit() bool {
	b1Count := float64(s.b1.count)
	b2Count := float64(s.b2.count)
	combinedCount := b1Count + b2Count

	if b1Count == 0 {
		return b2Count >= float64(s.capacity)
	}

	if combinedCount == 0 {
		return false
	}

	prevWindowWeight := b1Count / combinedCount
	approxReqCount := int(b1Count*prevWindowWeight) + int(b2Count)

	return approxReqCount >= s.capacity
}
