package minStack

import "math"

type Item struct {
	num int
	min int
}

type MinStack struct {
	stack []Item
}

func Constructor() MinStack {
	return MinStack{
		stack: []Item{},
	}
}

func (ms *MinStack) Push(val int) {
	var currentMin int
	if len(ms.stack) == 0 {
		currentMin = val
	} else {
		currentMin = ms.stack[len(ms.stack)-1].min
	}

	min := int(math.Min(float64(currentMin), float64(val)))
	ms.stack = append(ms.stack, Item{val, min})
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
}

func (ms *MinStack) Top() int {
	item := ms.stack[len(ms.stack)-1]
	return item.num
}

func (ms *MinStack) GetMin() int {
	item := ms.stack[len(ms.stack)-1]
	return item.min
}
