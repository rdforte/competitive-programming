package test

import (
	twoSum "github.com/rdforte/competitive-programming/LeetCode/Easy/1-two-sum/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tableTest := []struct {
		nums         []int
		target       int
		expectResult []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tst := range tableTest {
		got := twoSum.TwoSum(tst.nums, tst.target)
		assert.ElementsMatch(t, tst.expectResult, got)
	}
}
