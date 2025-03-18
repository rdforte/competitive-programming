package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(getResults([][]int{
		{1, 2},
		{2, 3, 3},
		{2, 3, 1},
		{2, 2, 2},
	}))

	fmt.Println(getResults([][]int{
		{1, 7},
		{2, 7, 6},
		{1, 2},
		{2, 7, 5},
		{2, 7, 6},
	}))

	fmt.Println(getResults([][]int{
		{2, 3, 1},
	}))

	fmt.Println(getResults([][]int{
		{1, 3},
		{2, 4, 2},
	}))

	fmt.Println(getResults([][]int{
		{1, 1},
		{1, 11},
		{1, 4},
		{1, 8},
		{2, 13, 7},
	}))
}

func getResults(queries [][]int) []bool {
	n := min(5e4, len(queries)*3) + 1
	obstaclePointSet := make(map[int]int)
	var points []int

	firstPoint, lastPoint := 0, n-1
	points = append(points, firstPoint, lastPoint)
	obstaclePointSet[firstPoint], obstaclePointSet[lastPoint] = 0, 0

	for _, q := range queries {
		typ := q[0]

		if typ == 1 {
			x := q[1]
			points = append(points, x)
		}
	}

	f := NewMaxFenwickTree(n)

	sort.Ints(points)

	// Add all points
	for i, p := range points {
		obstaclePointSet[p] = i // track index for point

		if i == 0 {
			continue
		}

		prevP := points[i-1]
		gapSize := p - prevP

		f.Add(p, gapSize)
	}

	var res []bool
	for q := len(queries) - 1; q >= 0; q-- {
		query := queries[q]
		typ := query[0]

		if typ == 1 {
			x := query[1]

			pointIdx := obstaclePointSet[x]
			prevPoint := points[pointIdx-1]
			nextPoint := points[pointIdx+1]

			distance := nextPoint - prevPoint

			// remove point from points
			// update index's for all points in struct that are before x.
			for k := range obstaclePointSet {
				if k == 0 || k < x {
					continue
				}
				obstaclePointSet[k]--
			}
			points = append(points[:pointIdx], points[pointIdx+1:]...)
			// delete(obstaclePointSet, x)

			f.Add(nextPoint, distance)
		}

		if typ == 2 {
			x := query[1]
			length := query[2]

			prevPoint := 0
			for _, p := range points {
				if p > x {
					break
				}
				prevPoint = p
			}

			distance := x - prevPoint
			fMax := f.Max(prevPoint)

			res = append(res, length <= max(distance, fMax))
		}

	}

	slices.Reverse(res)
	return res
}

type FenwickTree []int

func NewMaxFenwickTree(size int) FenwickTree {
	return make(FenwickTree, size)
}

func (f FenwickTree) Add(pos, val int) {
	for ; pos < len(f); pos |= (pos + 1) {
		f[pos] = max(f[pos], val)
	}
}

func (f FenwickTree) Max(pos int) (mx int) {
	for ; pos >= 0; pos = (pos & (pos + 1)) - 1 {
		mx = max(f[pos], mx)
	}
	return
}
