package main

import (
	"fmt"
	"math"
)

func main() {
	right2 := TreeNode{7, nil, nil}
	left2 := TreeNode{15, nil, nil}
	right1 := TreeNode{20, &left2, &right2}
	left1 := TreeNode{9, nil, nil}
	root := TreeNode{3, &left1, &right1}
	// left3 := TreeNode{4, nil, nil}
	// right3 := TreeNode{4, nil, nil}
	// left2 := TreeNode{3, &left3, &right3}
	// right2 := TreeNode{3, nil, nil}
	// right1 := TreeNode{2, &left2, &right2}
	// left1 := TreeNode{2, nil, nil}
	// root := TreeNode{1, &left1, &right1}

	fmt.Println(isBalanced(&root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balanced := true

	var heightOfSubtree func(n *TreeNode) int
	heightOfSubtree = func(n *TreeNode) int {
		if n == nil {
			return -1
		}

		leftHeight := heightOfSubtree(n.Left) + 1
		rightHeight := heightOfSubtree(n.Right) + 1

		if math.Abs(float64(leftHeight)-float64(rightHeight)) > 1 {
			balanced = false
		}

		return int(math.Max(float64(leftHeight), float64(rightHeight)))
	}

	heightOfSubtree(root)

	return balanced
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
