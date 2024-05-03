package main

import "fmt"

func main() {
	six := &TreeNode{6, nil, nil}
	four := &TreeNode{4, nil, nil}
	one := &TreeNode{1, nil, nil}
	two := &TreeNode{2, one, nil}
	three := &TreeNode{3, two, four}
	root := &TreeNode{5, three, six}

	fmt.Println(kthSmallest(root, 1))
	fmt.Println(kthSmallest(root, 2))
	fmt.Println(kthSmallest(root, 3))
	fmt.Println(kthSmallest(root, 4))
	fmt.Println(kthSmallest(root, 5))
	fmt.Println(kthSmallest(root, 6))
}

func kthSmallest(root *TreeNode, k int) int {
	var nodes []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		nodes = append(nodes, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	return nodes[k-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
