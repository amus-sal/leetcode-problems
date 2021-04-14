/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var levels [][]int

func levelOrder(root *TreeNode) [][]int {
	levels = make([][]int, 0)
	appendLevel(root, 0)
	return levels
}

func appendLevel(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(levels) <= level {
		newLevel := make([]int, 0)
		levels = append(levels, newLevel)
	}
	levels[level] = append(levels[level], node.Val)

	appendLevel(node.Left, level+1)
	appendLevel(node.Right, level+1)
}