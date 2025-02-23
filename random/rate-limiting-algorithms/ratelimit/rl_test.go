package ratelimit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rdforte/competitive-programming/random/rate-limiting-algorithms/ratelimit"
)

func TestRateLimiter_AddRequest_LimitsRequestsInSingleBucket(t *testing.T) {
	limit := 1
	interval := 10
	rl := ratelimit.New(interval, limit)

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	gotRateLimited := rl.AddRequest(date)
	gotRateLimited = rl.AddRequest(date)

	wantRateLimited := true
	assert.Equal(t, wantRateLimited, gotRateLimited)
}

func TestRateLimiter_AddRequest_LimitsRequestsIn2Buckets(t *testing.T) {
	limit := 1
	interval := 10
	rl := ratelimit.New(interval, limit)

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	gotRateLimited := rl.AddRequest(date)
	gotRateLimited = rl.AddRequest(date.Add(10 * time.Second))

	wantRateLimited := true
	assert.Equal(t, wantRateLimited, gotRateLimited)
}

func TestRateLimiter_AddRequest_LimitsRequestsNewIntervalResetsRateLimit(t *testing.T) {
	limit := 2
	interval := 10
	rl := ratelimit.New(interval, limit)

	date := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	gotRateLimited := rl.AddRequest(date)
	gotRateLimited = rl.AddRequest(date)

	wantRateLimited := true
	assert.Equal(t, wantRateLimited, gotRateLimited)

	gotRateLimited = rl.AddRequest(date.Add(11 * time.Second))

	wantRateLimited = false
	assert.Equal(t, wantRateLimited, gotRateLimited)
}
