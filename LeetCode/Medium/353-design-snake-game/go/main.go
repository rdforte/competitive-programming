package main

import (
	"fmt"
)

func main() {
	s := Constructor(3, 2, [][]int{{1, 2}, {0, 1}})
	fmt.Println(s.Move("R"))
	fmt.Println(s.Move("D"))
	fmt.Println(s.Move("R"))
	fmt.Println(s.Move("U"))
	fmt.Println(s.Move("L"))
	fmt.Println(s.Move("U"))

	fmt.Println("\nNEW GAME ------------------------")
	s = Constructor(2, 2, [][]int{{0, 1}})
	fmt.Println(s.Move("R"))
	fmt.Println(s.Move("D"))
}

var dirs = map[string][]int{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

type SnakeGame struct {
	height   int
	width    int
	snakeBod map[[2]int]struct{}
	food     [][]int
	tail     *Body
	head     *Body
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	grid := make([][]int, 0, height)
	for range height {
		grid = append(grid, make([]int, width))
	}

	grid[0][0] = 1

	dummyTail := &Body{}
	dummyHead := &Body{}
	head := &Body{
		prev: dummyTail,
		next: dummyHead,
	}
	dummyTail.next = head
	dummyHead.prev = head

	snakeBod := make(map[[2]int]struct{})
	snakeBod[[2]int{head.row, head.col}] = struct{}{}

	return SnakeGame{
		height:   height,
		width:    width,
		snakeBod: snakeBod,
		food:     food,
		tail:     dummyTail,
		head:     dummyHead,
	}
}

func (s *SnakeGame) Move(direction string) int {
	head := &Body{
		row: s.head.prev.row + dirs[direction][0],
		col: s.head.prev.col + dirs[direction][1],
	}

	if head.row < 0 || head.row >= s.height || head.col < 0 || head.col >= s.width {
		return -1
	}

	// eats food
	if len(s.food) > 0 && s.food[0][0] == head.row && s.food[0][1] == head.col {
		s.food = s.food[1:]
		s.addHead(head)

		return len(s.snakeBod) - 1
	}

	tail := s.tail.next
	delete(s.snakeBod, [2]int{tail.row, tail.col})

	tail.prev.next = tail.next
	tail.next.prev = tail.prev

	// hits own body
	if _, ok := s.snakeBod[[2]int{head.row, head.col}]; ok {
		return -1
	}

	s.addHead(head)

	return len(s.snakeBod) - 1
}

func (s *SnakeGame) addHead(b *Body) {
	s.snakeBod[[2]int{b.row, b.col}] = struct{}{}
	s.head.prev.next = b
	b.prev = s.head.prev
	s.head.prev = b
	b.next = s.head
}

type Body struct {
	row  int
	col  int
	next *Body
	prev *Body
}
