/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var concat string
	concat = ""
	return dfs(root, concat)
}

func dfs(root *TreeNode, concat string) int {
	if root == nil {
		return 0
	}
	concat = fmt.Sprintf("%v%v", concat, root.Val)
	if root.Left == nil && root.Right == nil {
		num, _ := strconv.Atoi(concat)
		return num
	}

	return dfs(root.Left, concat) + dfs(root.Right, concat)
}
