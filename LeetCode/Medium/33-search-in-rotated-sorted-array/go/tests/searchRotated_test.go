package searchRotated_test

import (
	"testing"

	searchRotated "github.com/rdforte/competitive-programming/LeetCode/Medium/33-search-in-rotated-sorted-array/go"
	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	t.Run("finds target when target is pivot", func(t *testing.T) {
		got := searchRotated.Search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
		want := 4
		assert.Equal(t, want, got)
	})

	t.Run("finds target when target is right of pivot", func(t *testing.T) {
		got := searchRotated.Search([]int{4, 5, 6, 7, 0, 1, 2}, 2)
		want := 6
		assert.Equal(t, want, got)
	})

	t.Run("finds target when target is left of pivot", func(t *testing.T) {
		got := searchRotated.Search([]int{4, 5, 6, 7, 0, 1, 2}, 5)
		want := 1
		assert.Equal(t, want, got)
	})

	t.Run("finds target when target is middle and not 0", func(t *testing.T) {
		got := searchRotated.Search([]int{5, 1, 3}, 5)
		want := 0
		assert.Equal(t, want, got)
	})

	t.Run("finds target when target is left and pivot and nums are not consecutive", func(t *testing.T) {
		got := searchRotated.Search([]int{8, 9, 2, 3, 4}, 9)
		want := 1
		assert.Equal(t, want, got)
	})
}
