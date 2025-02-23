package ratelimit

import (
	"math"
	"sync"
	"time"
)

type RateLimiter struct {
	limit    int
	interval time.Duration
	b1       *Bucket
	b2       *Bucket
	mu       sync.Mutex
}

type Bucket struct {
	startTime time.Time
	count     int
}

func New(interval, limit int) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		interval: time.Duration(interval) * time.Second,
		b1:       &Bucket{},
		b2:       &Bucket{},
	}
}

func (rl *RateLimiter) AddRequest(now time.Time) bool {
	startTime := now.Truncate(rl.interval)

	rl.mu.Lock()
	{
		if startTime.Equal(rl.b2.startTime.Add(rl.interval)) {
			rl.b1 = rl.b2
			rl.b2 = &Bucket{
				startTime: startTime,
			}
		}

		if startTime.Add(-rl.interval).After(rl.b2.startTime) {
			rl.b1 = &Bucket{
				startTime: startTime.Add(-rl.interval * 2),
			}
			rl.b2 = &Bucket{
				startTime: startTime.Add(-rl.interval),
			}
		}

		rl.b2.count = min(rl.b2.count, math.MaxInt-1) + 1
	}
	rl.mu.Unlock()

	b1Count := int(rl.interval-(now.Sub(startTime)*time.Second)) * rl.b1.count

	return b1Count+rl.b2.count >= rl.limit
}
