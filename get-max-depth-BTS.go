/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	return getDepth(root)
}
func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return max(getDepth(node.Left), getDepth(node.Right)) + 1
}
func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}