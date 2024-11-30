package main

import "fmt"

func main() {
	s := Constructor()
	s.Push(5)
	s.Push(7)
	s.Push(5)
	s.Push(7)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

type FreqStack struct {
	numFrequency   map[int]int
	frequencyStack map[int][]int
	maxFrequency   int
}

func Constructor() FreqStack {
	return FreqStack{
		numFrequency:   make(map[int]int),
		frequencyStack: make(map[int][]int),
		maxFrequency:   0,
	}
}

func (this *FreqStack) Push(val int) {
	this.numFrequency[val]++
	frequency := this.numFrequency[val]
	this.maxFrequency = max(this.maxFrequency, frequency)
	this.frequencyStack[frequency] = append(this.frequencyStack[frequency], val)
}

func (this *FreqStack) Pop() int {
	fIdx := this.maxFrequency
	stackLen := len(this.frequencyStack[fIdx])
	num := this.frequencyStack[fIdx][stackLen-1]
	this.frequencyStack[fIdx] = this.frequencyStack[fIdx][:stackLen-1]

	this.numFrequency[num]--

	if len(this.frequencyStack[fIdx]) == 0 {
		this.maxFrequency--
	}

	return num
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
