package visitRooms_test

import (
	"testing"

	visitRooms "github.com/rdforte/competitive-programming/LeetCode/Medium/841-keys-and-rooms/go"
	"github.com/stretchr/testify/assert"
)

func TestCanVisitAllRooms(t *testing.T) {
	t.Run("should visit all rooms", func(t *testing.T) {
		got := visitRooms.CanVisitAllRooms([][]int{{1}, {2}, {3}, {}})
		want := true
		assert.Equal(t, want, got)
	})

	t.Run("should not visit all rooms", func(t *testing.T) {
		got := visitRooms.CanVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}})
		want := false
		assert.Equal(t, want, got)
	})
}
