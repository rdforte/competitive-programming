package happyNubmer_test

import (
	"testing"

	happyNumber "github.com/rdforte/competitive-programming/LeetCode/Easy/202-happy-number/go"
	"github.com/stretchr/testify/assert"
)

func TestHappyNumber(t *testing.T) {
	t.Run("should return true if n is happy number", func(t *testing.T) {
		isHappy := happyNumber.Ishappy(19)
		assert.Equal(t, true, isHappy)
	})

	t.Run("should return false when n is not a happy number", func(t *testing.T) {
		isHappy := happyNumber.Ishappy(2)
		assert.Equal(t, false, isHappy)
	})
}
