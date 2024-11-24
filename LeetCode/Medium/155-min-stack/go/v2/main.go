package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
}

type MinStack struct {
	stack   []int
	minNums []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minNums) == 0 {
		this.minNums = append(this.minNums, val)
		return
	}

	topMinNum := this.minNums[len(this.minNums)-1]
	if val <= topMinNum {
		this.minNums = append(this.minNums, val)
	}
}

func (this *MinStack) Pop() {
	topMin := this.minNums[len(this.minNums)-1]
	topStack := this.stack[len(this.stack)-1]

	if topMin == topStack {
		this.minNums = this.minNums[:len(this.minNums)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minNums[len(this.minNums)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
