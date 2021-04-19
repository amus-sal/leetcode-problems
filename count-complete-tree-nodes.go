/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	leftHeight := leftHeight(root)
	rightHeight := rightHeight(root)
	if leftHeight == rightHeight {
		return (1 << rightHeight) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func rightHeight(root *TreeNode) int {
	height := 0
	for root != nil {
		height++
		root = root.Right
	}
	return height
}

func leftHeight(root *TreeNode) int {
	height := 0
	for root != nil {
		height++
		root = root.Left
	}
	return height
}