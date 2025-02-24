package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	ratelimiter "github.com/rdforte/competitive-programming/at/rnd1"
)

func TestUserRateLimiter(t *testing.T) {
	userLimiter := ratelimiter.NewUserRateLimiter()

	userID := "mock-user"

	limit := 2
	interval := 5
	userLimiter.AddUser(userID, interval, limit)

	startTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	canPass := userLimiter.Limit(startTime, userID)
	assert.True(t, canPass)
	canPass = userLimiter.Limit(startTime, userID)
	assert.False(t, canPass)
}

func TestUserRateLimiter_SlideSingleWindow(t *testing.T) {
	userLimiter := ratelimiter.NewUserRateLimiter()

	userID := "mock-user"

	limit := 2
	interval := 5
	userLimiter.AddUser(userID, interval, limit)

	startTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	canPass := userLimiter.Limit(startTime, userID)
	assert.True(t, canPass)
	canPass = userLimiter.Limit(startTime.Add(5*time.Second), userID)
	assert.False(t, canPass)
}

func TestUserRateLimiter_SlideSingleWindowWithSeconds(t *testing.T) {
	userLimiter := ratelimiter.NewUserRateLimiter()

	userID := "mock-user"

	limit := 2
	interval := 5
	userLimiter.AddUser(userID, interval, limit)

	startTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	canPass := userLimiter.Limit(startTime, userID)
	assert.True(t, canPass)
	canPass = userLimiter.Limit(startTime.Add(9*time.Second), userID)
	assert.True(t, canPass)
}
