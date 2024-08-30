package main

import "fmt"

func main() {
	fmt.Println(minDays([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}) == 2)
	fmt.Println(minDays([][]int{
		{0, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
		{1, 1, 1, 1},
	}) == 0)
}

// RULES for Articulation Point (a node that when removed will seperate the graph in 2)
//
// 1. A non-root node is an articulation point if it has a child whose lowest reachable time is greater than or equal to the node's discovery time.
// This condition means that the child (and its subtree) cannot reach any ancestor (any node that lies on the path from the root node to the current node)
// of the current node without going through the current node, making it critical for connectivity.
//
// 2. The root node of the DFS tree is an articulation point if it has more than one child. Removing the root would disconnect these children from each other.

func minDays(grid [][]int) int {
	directions := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	landCells, numIslands := 0, 0
	apInfo := newArticulationPoint(false, 0)

	// setup slices to store info for each cell
	rows, cols := len(grid), len(grid[0])
	discoverTime, lowestReachableTime, parentCell := newCellInfo(rows, cols)

	var findArticulationPoint func(r, c int)
	findArticulationPoint = func(r, c int) {
		discoverTime[r][c] = apInfo.time
		apInfo.time++
		lowestReachableTime[r][c] = discoverTime[r][c]
		children := 0

		for _, dir := range directions {
			newRow, newCol := r+dir[0], c+dir[1]
			if isValidLandCell(newRow, newCol, grid) {
				if discoverTime[newRow][newCol] == -1 {
					children++
					parentCell[newRow][newCol] = r*cols + c // row index * num cols + col index = unique individual number for cell
					findArticulationPoint(newRow, newCol)

					// update lowest reachable time
					lowestReachableTime[r][c] = min(lowestReachableTime[r][c], lowestReachableTime[newRow][newCol])

					// check for articulation point condition
					if lowestReachableTime[newRow][newCol] >= discoverTime[r][c] &&
						parentCell[r][c] != -1 {
						apInfo.hasArticulationPoint = true
					}
				} else if newRow*cols+newCol != parentCell[r][c] {
					// update lowestReachableTime
					lowestReachableTime[r][c] = min(lowestReachableTime[r][c], discoverTime[newRow][newCol])
				}
			}
		}

		// root of dfs tree is an articulation point if it has more than one child.
		if parentCell[r][c] == -1 && children > 1 {
			apInfo.hasArticulationPoint = true
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				// found land cell
				landCells++
				if discoverTime[r][c] == -1 {
					// have not visited this land before
					// start dfs to find the new island
					findArticulationPoint(r, c)
					numIslands++
				}
			}
		}
	}

	// Determine minimum number of days to disconnect the grid
	if numIslands == 0 || numIslands >= 2 {
		return 0 // no islands or more than 2 islands so no days needed to divide land.
	}

	if landCells == 1 {
		return 1 // only one cell was land so therefore only 1 day needed to remove it.
	}

	// landCells is more than 1 and we have 1 land mass so how many days needed to divide land?

	if apInfo.hasArticulationPoint {
		return 1 // we have an articulation point meaning changing this cell to 0 would break the graph into 2 ie: 2 islands.
	}

	// There is no articulation point so atleast 2 days will be needed to break the land mass into 2 islands.
	// We will never need more than 2 days to break the graph (land mass) into 2.
	return 2
}

// HELPERS --------------------------------------------------------------------------------------

func newArticulationPoint(hasArticulationPoint bool, time int) ArticulationPoint {
	return ArticulationPoint{hasArticulationPoint, time}
}

type ArticulationPoint struct {
	hasArticulationPoint bool
	time                 int
}

// newCellInfo returns the 3 key pieces of information we need to keep track of in order to perform `Tarjan algorithm`.
// discoveryTime: When a node is first visited during DFS.
// lowestReachableTime: Minimum discoveryTime of any node that can be reached from the subtree rooted at the current node,
// including the current node itself.
// parentCell: Is the node from which the current node was discovered during DFS.
func newCellInfo(numRows, numCols int) (discoveryTime [][]int, lowestReachableTime [][]int, parentCell [][]int) {
	discoveryTime = make([][]int, 0, numRows)
	lowestReachableTime = make([][]int, 0, numRows)
	parentCell = make([][]int, 0, numRows)
	for range numRows {
		dt := make([]int, 0, numCols)
		lr := make([]int, 0, numCols)
		pc := make([]int, 0, numCols)
		for range numCols {
			dt = append(dt, -1)
			lr = append(lr, -1)
			pc = append(pc, -1)
		}
		discoveryTime = append(discoveryTime, dt)
		lowestReachableTime = append(lowestReachableTime, lr)
		parentCell = append(parentCell, pc)
	}
	return
}

func isValidLandCell(r, c int, grid [][]int) bool {
	numRows, numCols := len(grid), len(grid[0])
	return r >= 0 && r < numRows && c >= 0 && c < numCols && grid[r][c] == 1
}
