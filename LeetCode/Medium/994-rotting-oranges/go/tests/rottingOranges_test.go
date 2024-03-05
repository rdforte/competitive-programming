package rottingOranges_test

import (
	"testing"

	rottingOranges "github.com/rdforte/competitive-programming/LeetCode/Medium/994-rotting-oranges/go"
	"github.com/stretchr/testify/assert"
)

func TestRottingOranges(t *testing.T) {
	t.Run("should return minutes when oranges are 4 directionally adjacent", func(t *testing.T) {
		grid := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}

		got := rottingOranges.OrangesRotting(grid)
		want := 4

		assert.Equal(t, want, got)
	})

	t.Run("should return -1 when oranges are not 4 directionally adjacent", func(t *testing.T) {
		grid := [][]int{
			{2, 1, 0},
			{0, 1, 1},
			{1, 0, 1},
		}

		got := rottingOranges.OrangesRotting(grid)
		want := -1

		assert.Equal(t, want, got)
	})

	t.Run("should return 0 when all oranges are already rotten", func(t *testing.T) {
		grid := [][]int{
			{0, 2},
		}

		got := rottingOranges.OrangesRotting(grid)
		want := 0

		assert.Equal(t, want, got)
	})
}
