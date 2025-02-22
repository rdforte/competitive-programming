package ratelimit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rdforte/competitive-programming/random/rate-limiting-algorithms/ratelimit"
)

func TestRateLimit_SlidingWindow_ReturnsTrueWhenExceededRateLimitInSameWindow(t *testing.T) {
	interval := 10
	numRequests := 2

	slidingWindow := ratelimit.New(interval, numRequests)

	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	exceededRateLimit := slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime)

	assert.True(t, exceededRateLimit)
}

func TestRateLimit_SlidingWindow_ReturnsTrueWhenExceededRateLimitWhenSlideOneWindow(t *testing.T) {
	interval := 10
	numRequests := 2

	slidingWindow := ratelimit.New(interval, numRequests)

	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	exceededRateLimit := slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime.Add(time.Second * 10))
	exceededRateLimit = slidingWindow.RateLimit(startTime.Add(time.Second * 10))

	assert.True(t, exceededRateLimit)
}

func TestRateLimit_SlidingWindow_ReturnsTrueWhenExceededRateLimitWhenSlideTwoWindow(t *testing.T) {
	interval := 10
	numRequests := 2

	slidingWindow := ratelimit.New(interval, numRequests)

	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	exceededRateLimit := slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime.Add(time.Second * 20))
	exceededRateLimit = slidingWindow.RateLimit(startTime.Add(time.Second * 20))

	assert.True(t, exceededRateLimit)
}

func TestRateLimit_SlidingWindow_ReturnsFalseWhenDoesNotExceedNumRequests(t *testing.T) {
	interval := 10
	numRequests := 2

	slidingWindow := ratelimit.New(interval, numRequests)

	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	exceededRateLimit := slidingWindow.RateLimit(startTime)
	exceededRateLimit = slidingWindow.RateLimit(startTime.Add(time.Second * 15))

	assert.False(t, exceededRateLimit)
}
