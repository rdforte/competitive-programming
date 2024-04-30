package main

import "fmt"

func main() {
	right2 := &TreeNode{7, nil, nil}
	left2 := &TreeNode{15, nil, nil}
	right1 := &TreeNode{20, left2, right2}
	left1 := &TreeNode{9, nil, nil}
	root := &TreeNode{3, left1, right1}

	fmt.Println(levelOrder(root))
	fmt.Println(levelOrder(nil))
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode

	if root == nil {
		return res
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		var level []int
		for i := 0; i < qLen; i++ {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, level)
		queue = queue[qLen:]
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
