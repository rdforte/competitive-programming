package minStack_test

import (
	minStack "github.com/rdforte/competitive-programming/LeetCode/Medium/155-min-stack/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := minStack.Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	gotMin := stack.GetMin() // return -3
	assert.Equal(t, -3, gotMin)
	stack.Pop()
	stack.Top()             // return 0
	gotMin = stack.GetMin() // return -2
	assert.Equal(t, -2, gotMin)
}
