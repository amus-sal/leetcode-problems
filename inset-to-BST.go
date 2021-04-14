/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	checkNode(root, val)
	return root
}

func checkNode(node *TreeNode, val int) {
	if node != nil {
		if node.Right == nil && val > node.Val {
			node.Right = &TreeNode{val, nil, nil}
		} else if node.Left == nil && val < node.Val {
			node.Left = &TreeNode{val, nil, nil}
		} else {
			if val > node.Val {
				checkNode(node.Right, val)
			} else {
				checkNode(node.Left, val)
			}
		}
		return
	}

	return

}