package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(clearStars("aaba*") == "aab")
	fmt.Println(clearStars("abc") == "abc")
	fmt.Println(clearStars("abc*abc*") == "bcbc")
	fmt.Println(clearStars("abc*abac*") == "bcabc")
	fmt.Println(clearStars("d*o*") == "")
	fmt.Println(clearStars("dk**") == "")
	fmt.Println(clearStars("dkz***") == "")
}

func clearStars(s string) string {
	var mark byte = 'z' + 1
	b := []byte(s)
	h := &charHeap{}
	heap.Init(h)

	for i := 0; i < len(b); i++ {
		if b[i] == '*' {
			b[i] = mark
			top := heap.Pop(h).(Char)
			b[top.Index] = mark
		} else {
			heap.Push(h, Char{b[i], i})
		}
	}

	str := []byte{}
	for _, c := range b {
		if c != mark {
			str = append(str, c)
		}
	}

	return string(str)
}

type charHeap []Char

func (h charHeap) Len() int { return len(h) }
func (h charHeap) Less(i, j int) bool {
	if h[i].Char == h[j].Char {
		return h[i].Index > h[j].Index
	}
	return h[i].Char < h[j].Char
}
func (h charHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *charHeap) Push(x any) {
	*h = append(*h, x.(Char))
}

func (h *charHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Char struct {
	Char  byte
	Index int
}
