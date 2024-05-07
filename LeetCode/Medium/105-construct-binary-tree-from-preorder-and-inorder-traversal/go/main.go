package main

import "fmt"

func main() {

	tree := buildTree([]int{3, 9, 10, 11, 20, 15, 7}, []int{10, 9, 11, 3, 15, 20, 7})

	fmt.Println(tree)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pos := make(map[int]int)

	for i := range inorder {
		pos[inorder[i]] = i
	}

	p := 0

	var dfs func(l, r int) *TreeNode
	dfs = func(l, r int) *TreeNode {
		if l > r || p >= len(preorder) {
			return nil
		}

		node := &TreeNode{preorder[p], nil, nil}

		p++

		node.Left = dfs(l, pos[node.Val]-1)
		node.Right = dfs(pos[node.Val]+1, r)

		return node
	}

	return dfs(0, len(preorder)-1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
