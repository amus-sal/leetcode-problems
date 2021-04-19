/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
	out := 0
	sumBST(root, low, high, &out)
	return out
}

func sumBST(root *TreeNode, low int, high int, out *int) {
	if root == nil {
		return
	}
	if root.Val >= low && root.Val <= high {
		*(out) += root.Val
	}
	sumBST(root.Left, low, high, out)
	sumBST(root.Right, low, high, out)
}