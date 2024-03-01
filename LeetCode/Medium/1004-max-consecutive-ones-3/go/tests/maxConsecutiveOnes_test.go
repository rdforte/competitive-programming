package maxConsecutiveOnes_test

import (
	"testing"

	maxOnes "github.com/rdforte/competitive-programming/LeetCode/Medium/1004-max-consecutive-ones-3/go"
	"github.com/stretchr/testify/assert"
)

func TestMaxConsecutiveOnes(t *testing.T) {
	t.Run("should return the maximum number of consecutive 1's in the array if you can flip at most k 0's.", func(t *testing.T) {
		table := []struct {
			k        int
			nums     []int
			expected int
		}{
			{2, []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 6},
			{3, []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 10},
		}

		for _, tt := range table {
			output := maxOnes.LongestOnes(tt.nums, tt.k)
			assert.Equal(t, tt.expected, output)
		}
	})
}
