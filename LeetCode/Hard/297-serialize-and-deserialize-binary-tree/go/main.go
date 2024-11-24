package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	five := TreeNode{Val: 5}
	four := TreeNode{Val: 4}
	three := TreeNode{Val: 3, Left: &four, Right: &five}
	two := TreeNode{Val: 2}
	root := TreeNode{Val: 1, Left: &two, Right: &three}

	c := Constructor()
	str := c.serialize(&root)
	res := c.deserialize(str)
	fmt.Println(res)
}

// Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer,
// or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var str []string
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			str = append(str, "*")
			return
		}
		str = append(str, strconv.Itoa(n.Val))
		dfs(n.Left)
		dfs(n.Right)
	}

	dfs(root)

	return strings.Join(str, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")
	idx := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if str[idx] == "*" {
			return nil
		}

		num, _ := strconv.Atoi(str[idx])
		node := &TreeNode{Val: num}
		idx++
		node.Left = dfs()
		idx++
		node.Right = dfs()

		return node
	}

	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
