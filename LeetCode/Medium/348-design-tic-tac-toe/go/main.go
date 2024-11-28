package main

import "fmt"

func main() {
	fmt.Println("GAME 1 ---------------")
	game1 := Constructor(3)
	fmt.Println(game1.Move(0, 0, 1))
	fmt.Println(game1.Move(0, 2, 2))
	fmt.Println(game1.Move(2, 2, 1))
	fmt.Println(game1.Move(1, 1, 2))
	fmt.Println(game1.Move(2, 0, 1))
	fmt.Println(game1.Move(1, 0, 2))
	fmt.Println(game1.Move(2, 1, 1))
	fmt.Println("GAME 2 ---------------")
	game2 := Constructor(2)
	fmt.Println(game2.Move(0, 0, 2))
	fmt.Println(game2.Move(1, 1, 1))
	fmt.Println(game2.Move(0, 1, 2))
}

type TicTacToe struct {
	players  []Player
	GridSize int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		players:  []Player{NewPlayer(n), NewPlayer(n)},
		GridSize: n,
	}
}

// Time Complexity = O(1)
// constant time to update verticals, horizontals, diagonals.
// constant time to determine if has won by vertical, horizontal or diagonal.
func (this *TicTacToe) Move(row int, col int, player int) int {
	p := player - 1
	if isVTB(row, col) {
		this.players[p].VTB++
	}
	if isVBT(row, col, this.GridSize) {
		this.players[p].VBT++
	}
	this.players[p].H[row]++
	this.players[p].V[col]++

	wonVertical := this.players[p].V[col] == this.GridSize
	wonHorizontal := this.players[p].H[row] == this.GridSize
	wonVTB := this.players[p].VTB == this.GridSize
	wonVBT := this.players[p].VBT == this.GridSize

	isWinner := wonVertical || wonHorizontal || wonVTB || wonVBT

	if isWinner {
		return player
	}

	return 0
}

func isVTB(row, col int) bool {
	return row == col
}

func isVBT(row, col, gridSize int) bool {
	return row+col == gridSize-1
}

func NewPlayer(gridSize int) Player {
	return Player{
		H: make([]int, gridSize),
		V: make([]int, gridSize),
	}
}

// Player represents a TicTacToe player.
// Space - O(n) where n is the grid size.
type Player struct {
	// H - horizontal.
	H []int
	// V - vertical.
	V []int
	// VTB - vertical top to bottom.
	VTB int
	// VBT - vertical bottom to top.
	VBT int
}
