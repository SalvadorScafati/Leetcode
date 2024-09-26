/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var same *bool
	value := true
	same = &value
	sameTree(p, q, same)
	return *same
}

func sameTree(p, q *TreeNode, same *bool) {
	if p == nil && q == nil {
		if *same {
			*same = true
			return
		}
	}
	if p == nil || q == nil {
		*same = false
		return
	}

	if p.Val != q.Val {
		*same = false
		return
	}

	sameTree(p.Left, q.Left, same)
	sameTree(p.Right, q.Right, same)
}
