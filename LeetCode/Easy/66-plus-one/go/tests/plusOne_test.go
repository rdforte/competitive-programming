package plusOne_test

import (
	"testing"

	plusOne "github.com/rdforte/competitive-programming/LeetCode/Easy/66-plus-one/go"
	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	t.Run("should increment last digit by 1", func(t *testing.T) {
		got := plusOne.PlusOne([]int{1, 2, 3})
		want := []int{1, 2, 4}
		assert.Equal(t, want, got)
	})

	t.Run("should increment last digit by 1 and add remainder to start", func(t *testing.T) {
		got := plusOne.PlusOne([]int{9})
		want := []int{1, 0}
		assert.Equal(t, want, got)
	})

	t.Run("should flip 999 to 1000", func(t *testing.T) {
		got := plusOne.PlusOne([]int{9, 9, 9})
		want := []int{1, 0, 0, 0}
		assert.Equal(t, want, got)
	})
}
