package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("V1 -----------")
	fmt.Println(checkValid([][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}))
	fmt.Println(checkValid([][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}))
	fmt.Println(checkValid([][]int{
		{1, 1},
		{1, 1},
	}))

	fmt.Println("V2 -----------")
	fmt.Println(checkValidV2([][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}))
	fmt.Println(checkValidV2([][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}))
	fmt.Println(checkValidV2([][]int{
		{1, 1},
		{1, 1},
	}))

	fmt.Println("V3 -----------")
	fmt.Println(checkValidV3([][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}))
	fmt.Println(checkValidV3([][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}))
	fmt.Println(checkValidV3([][]int{
		{1, 1},
		{1, 1},
	}))
	fmt.Println(checkValidV3([][]int{
		{1, 2},
		{2, 2},
	}))
}

// BitSet
func checkValidV3(matrix [][]int) bool {
	n := len(matrix)

	cmp := big.NewInt(1)
	for range n - 1 {
		cmp.Lsh(cmp, 1)
		cmp.Or(cmp, big.NewInt(1))
	}

	for i := range n {
		var row big.Int
		var col big.Int

		for j := range n {
			rNum := matrix[i][j]
			cNum := matrix[j][i]
			row.SetBit(&row, rNum-1, 1)
			col.SetBit(&col, cNum-1, 1)
		}

		if row.Cmp(cmp) != 0 || col.Cmp(cmp) != 0 {
			return false
		}
	}

	return true
}

func checkValidV2(matrix [][]int) bool {
	n := len(matrix)
	for i := range n {
		row := make(map[int]struct{}, n)
		col := make(map[int]struct{}, n)
		for j := range n {
			rNum := matrix[i][j]
			cNum := matrix[j][i]
			row[rNum] = struct{}{}
			col[cNum] = struct{}{}
		}

		if len(row) < n || len(col) < n {
			return false
		}
	}

	return true
}

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	rows := make(map[int][]bool, n)
	cols := make(map[int][]bool, n)

	for i := range n {
		r := make([]bool, n+1)
		c := make([]bool, n+1)
		r[0] = true
		c[0] = true

		rows[i] = r
		cols[i] = c
	}

	for r, rVal := range matrix {
		for c, cVal := range rVal {
			rows[r][cVal] = true
			cols[c][cVal] = true
		}
	}

	for i := range n {
		for v := range n + 1 {
			if !rows[i][v] {
				return false
			}
			if !cols[i][v] {
				return false
			}
		}
	}

	return true
}
