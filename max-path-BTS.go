/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var out int

func maxPathSum(root *TreeNode) int {
	out = -100000000
	output := getPath(root)
	out = max(out, output)

	return out
}

func getPath(node *TreeNode) int {

	if node == nil {
		return -10000000
	}
	rightPath := getPath(node.Right)
	leftPath := getPath(node.Left)
	out = max(out, node.Val)

	if leftPath != -10000000 && rightPath != -10000000 {
		out = max(out, rightPath+node.Val)
		out = max(out, leftPath+node.Val)
		out = max(out, rightPath+leftPath+node.Val)

		return max(max(rightPath, leftPath)+node.Val, node.Val)
	} else if leftPath != -10000000 {
		out = max(out, leftPath+node.Val)
		return max(leftPath+node.Val, node.Val)
	} else if rightPath != -10000000 {
		out = max(out, rightPath+node.Val)
		return max(rightPath+node.Val, node.Val)
	} else {
		return node.Val
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}