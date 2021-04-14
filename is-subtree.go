func isSubtree(s *TreeNode, t *TreeNode) bool {
	return f(s, t)
}

// resolve child problem
func compare(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return compare(node1.Left, node2.Left) && compare(node1.Right, node2.Right)
}

func f(s *TreeNode, t *TreeNode) bool {
	if compare(s, t) {
		return true
	}

	if s == nil {
		return false
	}

	// find all child problem
	if f(s.Left, t) || f(s.Right, t) {
		return true
	}

	return false
}