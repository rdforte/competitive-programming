package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(maxProbability(3, [][]int{
		{0, 1},
		{1, 2},
		{0, 2},
	}, []float64{0.5, 0.5, 0.2}, 0, 2) == 0.25)
}

type Node struct {
	id   int
	prob float64
}

// m = edges, n = nodes
// priority queue time complexity = O(log.m)
// if i have 3 nodes and 3 edges i have a total of 6 possibly Nodes im adding to my priority queue
// therefore (m+n).log.m
//
// space = O(m+n)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	adjList := make(map[int][]Node, n)
	probabilities := make([]float64, n)
	probabilities[start_node] = 1

	for i, e := range edges {
		n1 := e[0]
		n2 := e[1]
		prob := succProb[i]
		adjList[n1] = append(adjList[n1], Node{n2, prob})
		adjList[n2] = append(adjList[n2], Node{n1, prob})
	}

	h := &Heap{Node{start_node, 1}}
	heap.Init(h)

	for h.Len() > 0 {
		node, ok := heap.Pop(h).(Node)
		if !ok {
			panic("cant convert to Node")
		}

		if node.id == end_node {
			return node.prob
		}

		for _, n := range adjList[node.id] {
			n.prob *= node.prob
			// THIS CHECK IS IMPORTANT
			// we can get away with not having this check but then we are adding unecessary nodes
			// to our priority queue which then has to perform a log.n operation. This saves us from that.
			if n.prob > float64(probabilities[n.id]) {
				heap.Push(h, n)
				probabilities[n.id] = n.prob
			}
		}

		adjList[node.id] = nil
	}

	return 0
}

type Heap []Node

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool { return h[i].prob > h[j].prob }

func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
