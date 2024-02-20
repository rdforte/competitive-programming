package main

import "fmt"

// We utilise a binary tree and always fill left to right.
// the height of our tree will always be log n which is the √n. n being the number of elements in the tree.
// we can use an array to represent the tree. The root will be at index 0 and the left child of a node at index i will be at 2i + 1 and the right child will be at 2i + 2.
// an array allows us to easily find the parent or children of a node and esily swap nodes.

func main() {
	mh := MinHeap{}
	fmt.Println(mh.Top())
	mh.Insert(20)
	fmt.Println(mh.Top())
	mh.Insert(25)
	mh.Insert(26)
	mh.Insert(3)
	mh.Insert(15)
	mh.Insert(30)
	mh.Insert(1)
	fmt.Println("---------")
	fmt.Println(mh.Delete())
	fmt.Println(mh.Delete())
	fmt.Println(mh.Delete())
	fmt.Println(mh.Delete())
	fmt.Println(mh.Delete())
	fmt.Println(mh.Delete())
	fmt.Println(mh.Delete())
}

type MinHeap struct {
	data []int
}

func (h *MinHeap) Top() (int, error) {
	if len(h.data) == 0 {
		return -1, fmt.Errorf("There are 0 items in the heap")
	}
	return h.data[0], nil
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) Delete() (int, error) {
	if len(h.data) == 0 {
		return -1, fmt.Errorf("There are 0 items in the heap")
	}

	out := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)

	return out, nil
}

func (h *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h *MinHeap) leftChild(idx int) int {
	return (idx * 2) + 1
}

func (h *MinHeap) rightChild(idx int) int {
	return (idx * 2) + 2
}

// time = O(log n)
func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	parIdx := h.parent(idx)
	parVal := h.data[parIdx]
	curVal := h.data[idx]

	if parVal > curVal {
		h.data[parIdx], h.data[idx] = curVal, parVal
		h.heapifyUp(parIdx)
	}
}

// time = O(log n)
func (h *MinHeap) heapifyDown(idx int) {
	leftIdx := h.leftChild(idx)
	rightIdx := h.rightChild(idx)

	if idx >= len(h.data) || leftIdx >= len(h.data) {
		return
	}

	val := h.data[idx]
	leftVal := h.data[leftIdx]

	if rightIdx < len(h.data) {
		rightVal := h.data[rightIdx]
		// right is the smallest and we are greater than the smallest
		if leftVal > rightVal && val > rightVal {
			h.data[rightIdx], h.data[idx] = val, rightVal
			h.heapifyDown(rightIdx)
		} else if leftVal < rightVal && val > leftVal {
			h.data[leftIdx], h.data[idx] = val, leftVal
			h.heapifyDown(leftIdx)
		}
	} else if leftVal < val {
		// just compare to left val
		h.data[leftIdx], h.data[idx] = val, leftVal
		h.heapifyDown(leftIdx)
	}

}
