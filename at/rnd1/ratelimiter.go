/*
Atlassion rnd1
*/
package ratelimiter

import (
	"math"
	"time"
)

func NewUserRateLimiter() *UserRateLimiter {
	return &UserRateLimiter{
		users: make(map[string]*RateLimiter),
	}
}

type UserRateLimiter struct {
	users map[string]*RateLimiter
}

func (u *UserRateLimiter) AddUser(userID string, intervalSec, limit int) {
	u.users[userID] = NewRateLimiter(intervalSec, limit)
}

func (r *UserRateLimiter) Limit(now time.Time, userID string) bool {
	return r.users[userID].Limit(now)
}

type RateLimiter struct {
	interval      time.Duration
	limit         int
	prevBucket    *Bucket
	currentBucket *Bucket
}

type Bucket struct {
	startTime time.Time
	count     int
}

func NewRateLimiter(intervalSec, limit int) *RateLimiter {
	return &RateLimiter{
		interval:      time.Duration(intervalSec) * time.Second,
		limit:         limit,
		prevBucket:    &Bucket{},
		currentBucket: &Bucket{},
	}
}

func (r *RateLimiter) Limit(now time.Time) bool {
	startTime := now.Truncate(r.interval)

	if r.currentBucket.startTime.Add(r.interval) == startTime {
		r.prevBucket = r.currentBucket
		r.currentBucket.startTime = startTime
		r.currentBucket.count = 0
	}

	if startTime.Add(-r.interval).After(r.currentBucket.startTime) {
		r.prevBucket.startTime = startTime.Add(-r.interval)
		r.prevBucket.count = 0

		r.currentBucket.startTime = startTime
		r.currentBucket.count = 0
	}

	r.currentBucket.count = min(r.currentBucket.count, math.MaxInt-1) + 1

	durationInCurrentBucket := now.Sub(startTime)
	durationInPreviousBucket := r.interval - durationInCurrentBucket

	previouseBucketWeight := int(durationInPreviousBucket / r.interval)

	countOfPreviouseBucket := previouseBucketWeight * r.prevBucket.count

	totalCount := countOfPreviouseBucket + r.currentBucket.count

	return totalCount < r.limit
}
