package main

import (
	"fmt"
	"math"
)

func main() {
	five := &TreeNode{5, nil, nil}
	one2 := &TreeNode{1, nil, nil}
	three := &TreeNode{3, nil, nil}
	four := &TreeNode{4, one2, five}
	one1 := &TreeNode{1, three, nil}
	root := &TreeNode{3, one1, four}

	fmt.Println(goodNodes(root))
}

func goodNodes(root *TreeNode) int {

	res := 0

	var dfs func(node *TreeNode, max float64)
	dfs = func(node *TreeNode, max float64) {
		if node == nil {
			return
		}

		if float64(node.Val) >= max {
			res++
		}

		dfs(node.Left, math.Max(max, float64(node.Val)))
		dfs(node.Right, math.Max(max, float64(node.Val)))
	}

	dfs(root, math.Inf(-1))

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
