package main

import "fmt"

func main() {
	four := &TreeNode{4, nil, nil}
	five := &TreeNode{5, nil, nil}
	three := &TreeNode{3, nil, four}
	two := &TreeNode{2, nil, five}
	root := &TreeNode{1, two, three}

	fmt.Println(rightSideView(root))

	seven := &TreeNode{7, nil, nil}
	six := &TreeNode{6, nil, nil}
	four = &TreeNode{4, nil, nil}
	five = &TreeNode{5, six, seven}
	three = &TreeNode{3, nil, nil}
	two = &TreeNode{2, four, five}
	root = &TreeNode{1, two, three}

	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode) []int {
	var res []int
	var q []*TreeNode

	if root == nil {
		return res
	}

	q = append(q, root)

	for len(q) > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			if i == qLen-1 {
				res = append(res, q[i].Val)
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[qLen:]
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
