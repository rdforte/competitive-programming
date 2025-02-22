package ratelimit

import (
	"math"
	"time"
)

type SlidingWindow struct {
	interval    time.Duration
	numRequests int
	b1          Bucket
	b2          Bucket
}

type Bucket struct {
	startTime time.Time
	count     int
}

func New(intervalSec, numRequests int) *SlidingWindow {
	return &SlidingWindow{interval: time.Duration(intervalSec) * time.Second, numRequests: numRequests}
}

func (s *SlidingWindow) RateLimit(now time.Time) bool {
	startTime := now.Truncate(s.interval)

	if s.b2.startTime.Add(s.interval) == startTime {
		s.b1 = s.b2
		s.b2 = Bucket{startTime: startTime}
	}

	if s.b2.startTime.Add(s.interval).Before(startTime) {
		s.b1 = Bucket{startTime: startTime.Add(-s.interval)}
		s.b2 = Bucket{startTime: startTime}
	}

	s.b2.count = min(math.MaxInt, s.b2.count+1)

	timeInCurWindow := s.interval - now.Sub(startTime)
	overlap := timeInCurWindow / s.interval

	prevWindowCount := int(overlap) * s.b1.count

	return prevWindowCount+s.b2.count >= s.numRequests
}
