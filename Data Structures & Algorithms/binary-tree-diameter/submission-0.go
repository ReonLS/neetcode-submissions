/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    dia := 0

	//return longest connecting edge between two nodes
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int{
		if root == nil{
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		dia = max(dia, left + right)
		return 1 + max(left,right)
	}
	dfs(root)
	return dia
}
