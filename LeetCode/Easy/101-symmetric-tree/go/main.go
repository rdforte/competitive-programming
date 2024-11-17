package main

func main() {
}

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		qLen := len(queue)
		for l, r := 0, qLen-1; l < r; l, r = l+1, r-1 {
			if queue[l] == nil && queue[r] == nil {
				continue
			}
			if queue[l] == nil || queue[r] == nil {
				return false
			}
			if queue[l].Val != queue[r].Val {
				return false
			}
		}

		for i := 0; i < qLen; i++ {
			tn := queue[0]
			queue = queue[1:]

			if tn == nil {
				continue
			}

			queue = append(queue, tn.Left)
			queue = append(queue, tn.Right)
		}

	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
