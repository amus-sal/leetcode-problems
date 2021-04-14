/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func hasPathSum(root *TreeNode, sum int) bool {
	res = 0
	getTrg(root, 0, &sum)
	if res > 0 {
		return true
	}
	return false
}

func getTrg(root *TreeNode, sum int, trg *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		// is leaf node
		if sum+root.Val == *trg {
			res += 1
		}
		return
	}
	getTrg(root.Left, sum+root.Val, trg)
	getTrg(root.Right, sum+root.Val, trg)

}