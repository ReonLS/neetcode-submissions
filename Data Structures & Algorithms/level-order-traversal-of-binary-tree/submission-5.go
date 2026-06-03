/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    //BFS, Queue
    queue := []*TreeNode{root}
    final := [][]int{}

    //root empty
    if root == nil{
        return final
    }

    for len(queue) > 0{
        level := len(queue)
        nodes := []int{}

        for range level{
            node := queue[0]
            queue = queue[1:]

            nodes = append(nodes, node.Val)

            if node.Left != nil{
                queue = append(queue, node.Left)
            }
            if node.Right != nil{
                queue = append(queue, node.Right)
            }
        }
        final = append(final, nodes)
    }
    return final
}
