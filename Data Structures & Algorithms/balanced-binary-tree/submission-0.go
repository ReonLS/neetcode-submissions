/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    final := true

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int{
		if root == nil{
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		//condition buat ngecek, kalo differencenya more than 1
		if left-right > 1 || right-left > 1{
			final = false
		}
		return 1 + max(left,right)
	}
	dfs(root)
	return final
}
