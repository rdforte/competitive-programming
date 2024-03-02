package maze_test

import (
	"testing"

	nearestExit "github.com/rdforte/competitive-programming/LeetCode/Medium/1926-nearest-exit-from-entrance-in-maze/go"
	"github.com/stretchr/testify/assert"
)

func TestNearestExitFromEntranceInMaze(t *testing.T) {
	t.Run("should find the nearest exit in the maze", func(t *testing.T) {
		maze := [][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		}
		entrance := []int{1, 2}
		got := nearestExit.NearestExit(maze, entrance)
		want := 1
		assert.Equal(t, want, got)
	})

	t.Run("should find the nearest exit in the maze when the start position is at the edge", func(t *testing.T) {
		maze := [][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		}
		entrance := []int{1, 0}
		got := nearestExit.NearestExit(maze, entrance)
		want := 2
		assert.Equal(t, want, got)
	})
}
