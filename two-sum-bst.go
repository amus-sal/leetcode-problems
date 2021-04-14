/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var total map[int]int

func findTarget(root *TreeNode, k int) bool {
	total = make(map[int]int)
	out := get(root, k)
	fmt.Println(total)
	return out
}

func get(node *TreeNode, k int) bool {
	if _, found := total[node.Val]; found {
		return true
	}

	if _, found := total[k-node.Val]; !found {

		total[k-node.Val] = 1
		if node.Left != nil && get(node.Left, k) {
			return true
		}
		if node.Right != nil && get(node.Right, k) {
			return true
		}
		return false
	} else {
		return true
	}

}